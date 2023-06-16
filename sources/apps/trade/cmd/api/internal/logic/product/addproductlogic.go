package product

import (
	"context"
	"encoding/json"
	"mime/multipart"

	resPb "github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/trade/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddProductLogic struct {
	logx.Logger
	ctx        context.Context
	svcCtx     *svc.ServiceContext
	FileHeader multipart.FileHeader
	File       multipart.File
}

func NewAddProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddProductLogic {
	return &AddProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddProduct
//
//	@Description: 绑定商品信息
//	@receiver l
//	@param req
//	@return error
func (l *AddProductLogic) AddProduct(req *types.AddProductReq) error {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	// 查询商品是否上架 productBind && productBind
	uid, err := l.ctx.Value(xconst.JWT_USER_ID).(json.Number).Int64()
	productInfo, err := l.svcCtx.TradeRpc.GetProductByBindIdAndOwner(context.Background(), &pb.GetProductByBindIdAndOwnerReq{
		Uid:         uid,
		ProductBind: req.PriductBind,
	})
	if err != nil {
		l.Logger.WithFields(logx.Field("err", err)).Error("查询已上架商品失败")
		return xerr.NewErrMsg("上架商品失败")
	}
	if productInfo.Product != nil {
		return xerr.NewErrMsg("该商品已上架！请勿重复上架")
	}
	// 拉取资源信息匹配到商品
	res, err := l.svcCtx.ResourcesRpc.GetFileById(context.Background(), &resPb.GetFileByIdReq{ID: req.PriductBind})
	if err != nil {
		return xerr.NewErrMsg("上架商品失败")
	}
	if res.File == nil {
		return xerr.NewErrMsg("该商品绑定的资源已删除！不能上架已删除的商品")
	}
	if res.File.Owner != uid {
		return xerr.NewErrMsg("不能上架别人的商品")
	}

	if err != nil {
		return xerr.NewErrCode(xerr.SERVER_ERROR)
	}
	// 增加product
	_, err = l.svcCtx.TradeRpc.AddProduct(l.ctx, &pb.AddProductReq{
		Name:          req.Name,
		Type:          res.File.Class,
		ProductBind:   req.PriductBind,
		Owner:         uid,
		Price:         req.Price,
		ProductPoster: res.File.FilePoster,
	})
	if err != nil {
		return xerr.NewErrMsg("上架商品失败")
	}
	return nil
}
