package fileRes

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"

	orderPb "github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	tradePb "github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileResDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileResDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileResDataLogic {
	return &GetFileResDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetFileResData
//
//	@Description: 获取资源文件
//	@receiver l
//	@param req
//	@return []byte
//	@return string
//	@return error
func (l *GetFileResDataLogic) GetFileResData(req *types.DownLoadFileReq) ([]byte, string, error) {
	// 参数校验
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, "", xerr.NewErrMsg(validatorResult)
	}
	// 创建context
	deadlineCtx, cancelFunc := context.WithDeadline(context.Background(), utils.GetContextDefaultTime())
	defer cancelFunc()
	// 判断用户是否允许下载
	rpcData, err := l.svcCtx.ResourcesRpc.CheckDownloadAllow(deadlineCtx, &pb.CheckDownloadAllowReq{FileId: req.ResourceFileId})
	if err != nil {
		return nil, "", xerr.NewErrMsg("资源状态检测失败，请稍后再试！")
	}
	if !rpcData.IsAllow {
		return nil, "", xerr.NewErrMsg("该资源发布者不允许下载该资源，请等待其上架后购买，或者公开此资源！")
	}
	// 判断是否商品
	productData, err := l.svcCtx.TradeRpc.SearchProductByResourcesBind(deadlineCtx, &tradePb.SearchProductByResourcesBindReq{ResourceId: req.ResourceFileId})
	if err != nil {
		return nil, "", xerr.NewErrMsg("不存在该资源！")
	}
	if len(productData.ProductName) == 0 || productData.GetProductId() == 0 {
		// 不是商品，直接下载文件
		return l.DownLoadFile(req.ResourceFileId, productData.GetProductName(), deadlineCtx)
	} else {
		// 是商品，查询用户是否购买
		uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
		if err != nil {
			return nil, "", xerr.NewErrMsg("获取你的购买记录失败，请稍后再试")
		}
		order, err := l.svcCtx.OrderRpc.GetOrderInfoByUserIdAndProductId(deadlineCtx, &orderPb.GetOrderInfoByUserIdAndProductIdReq{
			ProductId: productData.ProductId,
			UserId:    uid,
		})
		if err != nil {
			return nil, "", xerr.NewErrMsg("获取你的购买记录失败，请稍后再试")
		}
		if order.Order == nil {
			return nil, "", xerr.NewErrMsg("您未购买该资源，请购买后在进行下载")
		}
		// 允许下载
		return l.DownLoadFile(req.ResourceFileId, productData.GetProductName(), deadlineCtx)
	}
}

// DownLoadFile
//
//	@Description: 使用商品id下载文件
//	@receiver l
//	@param resourcesId
//	@param productName
//	@param ctx
//	@return []byte
//	@return string
//	@return error
func (l *GetFileResDataLogic) DownLoadFile(resourcesId int64, productName string, ctx context.Context) ([]byte, string, error) {
	// 查询文件 key
	fileResInfo, err := l.svcCtx.ResourcesRpc.GetFileById(ctx, &pb.GetFileByIdReq{ID: resourcesId})
	if err == gorm.ErrRecordNotFound {
		return nil, "", nil
	}
	if err != nil {
		return nil, "", err
	}
	// 下载cos文件
	data, err := l.svcCtx.OSSClient.DownloadFile(fileResInfo.File.Link)
	if err != nil {
		return nil, "", err
	}
	return data, productName + fileResInfo.File.GetSuffix(), nil
}
