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

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login
// @Description: 登录函数
// @receiver l
// @param req
// @return resp
// @return err
func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// 校验请求参数
	if validatorResult := l.svcCtx.Validator.ValidateZh(req); len(validatorResult) > 0 {
		return nil, xerr.NewErrMsg(validatorResult)
	}
	// 校验验证码
	// if !l.svcCtx.Captcha.Verify(req.CaptchaId, req.Captcha, true) {
	// 	return nil, xerr.NewErrCode(xerr.CAPTCHA_CHECK_ERR)
	// }
	// 查询用户信息
	withDeadline, cancelFunc := context.WithDeadline(l.ctx, utils.GetContextDefaultTime())
	defer cancelFunc()
	userInfoReq := &pb.GetUserByPhoneOrEmailReq{
		Phone: req.Phone,
		Email: req.Email,
	}
	// rpc调用
	userData, err := l.svcCtx.UserRpc.GetUserByPhoneOrEmail(withDeadline, userInfoReq)
	if err != nil {
		return nil, err
	}
	// 检查密码
	if userData.User.Password != req.Password {
		return nil, xerr.NewErrCode(xerr.AUTH_CHECK_FAILURE)
	}
	// 回传用户数据token
	genTime := utils.GetJwtIatTime()
	expireGap := utils.GetJwtExpireDefaultTime()
	token, err := l.genJwtToken(l.svcCtx.Config.Auth.AccessSecret, genTime, expireGap,
		userData.User.UID, userData.User.UniqueID)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken: token,
		ExpireTime:  strconv.FormatInt(genTime+expireGap, 10),
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
func (l *LoginLogic) genJwtToken(secretKey string, iat, seconds, userId, userUniqueId int64) (string, error) {
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
