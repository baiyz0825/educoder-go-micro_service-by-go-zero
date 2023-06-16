package user

import (
	"context"
	"strconv"

	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/api/internal/types"
	"github.com/baiyz0825/school-share-buy-backend/apps/user/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/utils"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register
// @Description: 用户注册接口
// @receiver l
// @param req
// @return resp
// @return err
func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// 解析请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	// 校验验证吗
	// if !l.svcCtx.Captcha.Verify(req.CaptchaId, req.Captcha, true) {
	// 	return nil, xerr.NewErrCode(xerr.CAPTCHA_CHECK_ERR)
	// }
	// 成功 -> 入库
	// 创建用户对象
	user := &pb.AddUserReq{
		Name:     req.Username,
		Password: req.Password,
		Phone:    req.Phone,
	}
	deadlineCtx, cancelFunc := context.WithDeadline(l.ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	// check 用户是否存在，手机号唯一
	userOrigin, err := l.svcCtx.UserRpc.GetUserByPhoneOrEmail(deadlineCtx, &pb.GetUserByPhoneOrEmailReq{
		Phone: req.Phone,
	})
	if err != nil {
		return nil, err
	}
	if userOrigin.User != nil {
		return nil, xerr.NewErrMsg("手机号已经存在，请更换手机号注册")
	}
	// rpc调用
	_, err = l.svcCtx.UserRpc.AddUser(deadlineCtx, user)
	if err != nil {
		return nil, err
	}
	userInfoReq := &pb.GetUserByPhoneOrEmailReq{
		Phone: req.Phone,
	}
	// 获取注册用户信息
	userData, err := l.svcCtx.UserRpc.GetUserByPhoneOrEmail(deadlineCtx, userInfoReq)
	if err != nil {
		return nil, err
	}
	// 回传用户数据token
	genTime := utils.GetJwtIatTime()
	expireGap := utils.GetJwtExpireDefaultTime()
	token, err := l.genJwtToken(l.svcCtx.Config.Auth.AccessSecret, genTime, expireGap,
		userData.User.UID, userData.User.UniqueID)
	if err != nil {
		return nil, err
	}
	return &types.RegisterResp{
		AccessToken:  token,
		ExpireTime:   strconv.FormatInt(genTime+expireGap, 10),
		RefreshToken: "",
	}, nil
}

// genJwtToken
//
//	@Description: 生成对应jwt密钥
//	@receiver l
//	@param secretKey jwt密钥
//	@param iat 颁发时间
//	@param seconds 过期时间
//	@param userId 用户Id
//	@param userUniqueId 用户唯一id
//	@return string jwt串
//	@return error
func (l *RegisterLogic) genJwtToken(secretKey string, iat, seconds, userId, userUniqueId int64) (string, error) {
	claims := make(jwt.MapClaims)
	// 过期时间
	claims["exp"] = iat + seconds
	// jwt 颁发时间
	claims["iat"] = iat
	claims[xconst.JWT_USER_USERUNIQUEID] = userUniqueId
	claims[xconst.JWT_USER_ID] = userId
	// 生成jwt
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
