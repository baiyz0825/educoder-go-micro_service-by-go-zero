package fileRes

import (
	"context"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileResInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFileResInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileResInfoLogic {
	return &GetFileResInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// GetFileResInfo
//
//	@Description: 使用file id查询文件详情
//	@receiver l
//	@param req
//	@return resp
//	@return err
func (l *GetFileResInfoLogic) GetFileResInfo(req *types.FileResInfoReq) (resp *types.FileResInfoResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	// 查询数据
	data, err := l.svcCtx.ResourcesRpc.GetFileById(l.ctx, &pb.GetFileByIdReq{ID: req.FileResId})
	if err != nil || data.File == nil {
		return nil, err
	}
	resp = &types.FileResInfoResp{}
	err = copier.Copy(&resp.File, data.File)
	if err != nil {
		l.Logger.WithFields(logx.Field("err: ", err)).Error("发生copier错误")
		return nil, xerr.NewErrMsg("数据处理异常")
	}
	resp.File.FileType = data.File.Type
	return
}
