package user

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type ModInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewModInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModInfoLogic {
	return &ModInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ModInfo
// @Description: 修改用户个人数据
// @receiver l
// @param req
// @return error
func (l *ModInfoLogic) ModInfo(req *types.UserDataReq) error {
	// 校验请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return xerr.NewErrMsg(validatorResult)
	}
	userId, err := strconv.ParseInt(l.ctx.Value(xconst.JWT_USER_ID).(json.Number).String(), 10, 64)
	if err != nil {
		return err
	}
	userUniqueId, err := strconv.ParseInt(l.ctx.Value(xconst.JWT_USER_USERUNIQUEID).(json.Number).String(), 10, 64)
	if err != nil {
		return err
	}
	// rpc保存
	userInfo := &pb.UpdateUserReq{
		UID:      userId,
		UniqueID: userUniqueId,
		Name:     req.Username,
		Age:      req.Age,
		Gender:   int64(req.Gender),
		Phone:    req.Phone,
		Email:    req.Email,
		Grade:    req.Grade,
		Major:    req.Major,
		Sign:     req.Sign,
		Class:    req.Class,
	}
	ctx, cancelFunc := context.WithDeadline(l.ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	_, err = l.svcCtx.UserRpc.UpdateUser(ctx, userInfo)
	if err != nil {
		return err
	}
	return nil
}
