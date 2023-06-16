package user

import (
	"context"
	"encoding/json"
	"strconv"

	"golang.org/x/sync/errgroup"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UserInfo
// @Description: 查询个人用户数据
// @receiver l
// @param req
// @return resp
// @return err
func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	userId, err := strconv.ParseInt(l.ctx.Value(xconst.JWT_USER_ID).(json.Number).String(), 10, 64)
	if err != nil {
		return nil, err
	}
	idReq := &pb.GetUserByIdReq{ID: userId}
	ctx, cancelFunc := context.WithDeadline(l.ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	userInfo, err := l.svcCtx.UserRpc.GetUserById(ctx, idReq)
	if err != nil {
		return nil, err
	}
	// 保存部分返回值
	resp = &types.UserInfoResp{}
	resp.UserInfo.UserId = userInfo.User.UID
	resp.UserInfo.UserName = userInfo.User.Name
	resp.UserInfo.Sign = userInfo.User.Sign
	resp.UserInfo.UserUniqueId = userInfo.User.UniqueID
	resp.UserInfo.Avatar = userInfo.User.Avatar
	resp.UserInfo.Email = userInfo.User.Email
	resp.UserInfo.Phone = userInfo.User.Phone
	resp.UserInfo.Age = userInfo.User.Age
	resp.UserInfo.Grade = strconv.FormatInt(userInfo.User.Grade, 10)
	resp.UserInfo.Class = strconv.FormatInt(userInfo.User.Class, 10)
	// 保留评级后两位数字
	resp.UserInfo.Star = strconv.FormatFloat(userInfo.User.Star, 'f', 2, 64)
	majorId := userInfo.User.Major
	reqSocialBindUserId := userInfo.User.UID
	group, _ := errgroup.WithContext(ctx)

	// 初始化变量
	var majorData *pb.Major
	// 获取对应数据
	group.Go(func() error {
		// 查询 major rpc
		majorRpc := &pb.GetMajorByIdReq{ID: majorId}
		temp, err := l.svcCtx.UserRpc.GetMajorById(ctx, majorRpc)
		if err != nil {
			return err
		}
		majorData = temp.GetMajor()
		return nil
	})
	var thirds []*pb.ThirdBind
	group.Go(func() error {
		// 查询 social rpc
		thirdUserBind := &pb.GetThirdBindDataReq{
			UserID: reqSocialBindUserId,
		}
		thirdBindData, err := l.svcCtx.UserRpc.GetThirdBindData(ctx, thirdUserBind)
		if err != nil {
			return err
		}
		thirds = thirdBindData.Thirds
		return nil
	})
	// 阻塞等待
	err = group.Wait()
	if err != nil {
		return nil, err
	}
	// 后续数据处理
	if majorData != nil {
		resp.UserInfo.Major = majorData.Name
	}
	if thirds != nil {
		var bind []types.SocialBind
		for _, third := range thirds {
			temp := types.SocialBind{
				Id:   third.ThirdType,
				Name: third.ThirdName,
			}
			bind = append(bind, temp)
		}
		resp.UserInfo.SocialBind = bind
	}
	return resp, nil
}
