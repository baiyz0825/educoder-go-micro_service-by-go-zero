package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/baiyz0825/school-share-buy-backend/apps/mq/cmd/job/bo"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	tradePb "github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/hibiken/asynq"
	"github.com/smartwalle/alipay/v3"

	"github.com/zeromicro/go-zero/core/logx"
)

type DoOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoOrderLogic {
	return &DoOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DoOrder ----- 下订单
func (l *DoOrderLogic) DoOrder(in *pb.DoOrderReq) (*pb.DoOrderResp, error) {
	deadline, cancelFunc := context.WithTimeout(context.Background(), utils.GetContextDuration())
	defer cancelFunc()
	// 获取商品价格 productId
	product, err := l.svcCtx.TradeRpc.GetProductById(deadline, &tradePb.GetProductByIdReq{
		ID: in.ProductId,
	})
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).
			Error(fmt.Sprintf("用户(%v)下单，商品(id:%v)信息不存在", in.UserId, in.ProductId))
		return nil, xerr.NewErrMsg("该商品不存在，请重新选择")
	}
	// 创建订单数据（是否需要过期）
	// gen uuid && db data
	id, err := utils.GenSnowFlakeId()
	if err != nil {
		l.Logger.WithFields(logx.Field("error:", err)).Error("生成订单流水号失败")
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	// 加入db
	var orderData *model.Order
	if in.PayPath == xconst.PAY_PATH_ALIPAY {
		orderData = &model.Order{
			UUID:            id,
			ProductID:       product.Product.ID,
			SysModel:        xconst.ORDER_SYSTEM_MODE_TRADE,
			Status:          0,
			UserID:          in.GetUserId(),
			PayPrice:        &product.Product.Price,
			PayPath:         xconst.PAY_PATH_ALIPAY,
			PayPathOrderNum: strconv.FormatInt(id, 10),
		}
	} else {
		orderData = &model.Order{
			UUID:            id,
			ProductID:       product.Product.ID,
			SysModel:        xconst.ORDER_SYSTEM_MODE_TRADE,
			Status:          0,
			UserID:          in.GetUserId(),
			PayPrice:        &product.Product.Price,
			PayPath:         xconst.PAY_PATH_WECHAT,
			PayPathOrderNum: strconv.FormatInt(id, 10),
		}
	}
	orderQ := l.svcCtx.Query.Order
	err = orderQ.WithContext(deadline).Create(orderData)
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).
			Error(fmt.Sprintf("订单创建失败,商品(id:%v),用户(id:%v),支付渠道-0为微信 1为阿里 (%v)", product.Product.ID, in.UserId, in.PayPath))
		return nil, xerr.NewErrMsg("订单创建失败")
	}
	// 生成支付二维码
	trade := alipay.Trade{
		Subject:    product.Product.Name,
		OutTradeNo: strconv.FormatInt(id, 10),
		// 转为小数点后两位：带转化数字，浮点数，小数点后位数，字符串位数
		TotalAmount: strconv.FormatFloat(product.Product.Price, 'f', 2, 64),
		// 回调订单状态地址
		NotifyURL: l.svcCtx.Config.AliPay.PayNoticeCallBackUrl,
	}
	tradePreCreate := alipay.TradePreCreate{
		Trade:       trade,
		GoodsDetail: nil,
	}
	l.Logger.Debugf("生成的交易信息是：%v", trade)
	tradePreCreateResp, err := l.svcCtx.AliPayClient.TradePreCreate(tradePreCreate)
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).
			Error(fmt.Sprintf("订单创建失败,商品(id:%v),用户(id:%v),支付渠道-0为微信 1为阿里 (%v)", product.Product.ID, in.UserId, in.PayPath))
		return nil, xerr.NewErrMsg("订单创建失败")
	}
	if !tradePreCreateResp.IsSuccess() {
		return nil, xerr.NewErrMsg(fmt.Sprintf("支付宝，订单创建失败,支付宝异常:%v", tradePreCreateResp.Content.Code))
	}
	// 支付宝创建订单成功，更新db
	updateCondition := model.Order{Status: xconst.ORDER_STATUS_PAYING, PayCodeURL: &tradePreCreateResp.Content.QRCode}
	_, err = orderQ.WithContext(deadline).Where(orderQ.UUID.Eq(id)).Updates(updateCondition)
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).
			Error(fmt.Sprintf("订单状态更新失败,商品(id:%v),用户(id:%v),支付渠道-0为微信 1为阿里 (%v)", product.Product.ID, in.UserId, in.PayPath))
		return nil, xerr.NewErrMsg("订单创建失败")
	}
	// 加入mq 订单超时关闭，未超时，更新订单状态 TODO 这部分创建失败未控制支付宝以及数据库数据回滚
	mqJobStruct := &bo.OrderMqStruct{
		Uuid:            id,
		UserId:          in.UserId,
		PayPathOrderNum: strconv.FormatInt(id, 10),
	}
	jobsPayload, err := json.Marshal(mqJobStruct)
	if err != nil {
		l.Logger.WithFields(logx.Field("err:", err)).
			Error(fmt.Sprintf("设置定时关闭订单任务失败:序列化定时任务数据失败 %v商品(id:%v),用户(id:%v),支付渠道-0为微信 1为阿里 (%v)", err, product.Product.ID, in.UserId, in.PayPath))
		return nil, xerr.NewErrMsg("订单创建失败")
	}
	group := errgroup.Group{}
	group.Go(func() error {
		// 发送延时任务30min之后删除订单
		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(bo.DELETE_EXPIRE_JOBS, jobsPayload), asynq.ProcessIn(30*time.Minute))
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).
				Error(fmt.Sprintf("发送定时关闭订单任务失败: %v商品(id:%v),用户(id:%v),支付渠道-0为微信 1为阿里 (%v)", err, product.Product.ID, in.UserId, in.PayPath))
			return xerr.NewErrMsg("订单超时关闭任务提交失败")
		}
		return nil
	})
	// 加入订单状态查询
	group.Go(func() error {
		task := asynq.NewTask(bo.CHECK_ORDER_STATUS_JOBS, jobsPayload)
		_, err = l.svcCtx.AsynqClient.Enqueue(task, asynq.ProcessIn(20*time.Second))
		if err != nil {
			l.Logger.WithFields(logx.Field("err:", err)).
				Error(fmt.Sprintf("发送订单支付任务失败: %v商品(id:%v),用户(id:%v),支付渠道-0为微信 1为阿里 (%v)", err, product.Product.ID, in.UserId, in.PayPath))
			return xerr.NewErrMsg("订单支付状态检查任务提交失败")
		}
		return nil
	})
	err = group.Wait()
	if err != nil {
		// TODO 回滚
		l.Logger.WithFields(logx.Field("err:", err)).Error("用户订单创建失败，正在回滚订单")
		return nil, xerr.NewErrCode(xerr.SERVER_ERROR)
	}
	// 返回
	return &pb.DoOrderResp{
		Status:          xconst.ORDER_STATUS_PAYING,
		PayUrl:          tradePreCreateResp.Content.QRCode,
		PayPathOrderNum: tradePreCreateResp.Content.OutTradeNo,
	}, nil
}
