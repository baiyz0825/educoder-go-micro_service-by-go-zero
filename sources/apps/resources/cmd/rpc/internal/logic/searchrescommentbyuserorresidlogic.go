package logic

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchResCommentByUserOrResIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchResCommentByUserOrResIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchResCommentByUserOrResIdLogic {
	return &SearchResCommentByUserOrResIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchResCommentByUserOrResId
//
//	@Description: 查询某一个资源下评论 ， 查询用户全部评论
//	@receiver l
//	@param in
//	@return *pb.SearchResCommentByUserOrResIdResp
//	@return error
func (l *SearchResCommentByUserOrResIdLogic) SearchResCommentByUserOrResId(in *pb.SearchResCommentByUserOrResIdReq) (*pb.SearchResCommentByUserOrResIdResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	if in.GetOwner() == 0 && in.GetResourceID() == 0 {
		return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	}
	q := l.svcCtx.Query.ResComment
	var condition []gen.Condition
	query := q.WithContext(l.ctx)
	if in.GetOwner() != 0 && in.GetResourceID() != 0 {
		// 查询指定用户发表到指定资源id评论
		condition = append(condition, q.Owner.Eq(in.GetOwner()), q.ResourceID.Eq(in.GetResourceID()))
	} else if in.GetOwner() != 0 {
		// 查询指定用户全部评论
		condition = append(condition, q.Owner.Eq(in.GetOwner()))
	} else if in.GetResourceID() != 0 {
		// 查询指定资源评论
		condition = append(condition, q.ResourceID.Eq(in.GetResourceID()))
	}
	// 查询
	pageDatas, _, err := query.Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
	if err == gorm.ErrRecordNotFound {
		return &pb.SearchResCommentByUserOrResIdResp{}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
	}
	// copy
	var respData []*pb.ResComment
	for _, data := range pageDatas {
		temp := &pb.ResComment{}
		err := copier.Copy(temp, data)
		if err != nil {
			return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
		}
		temp.UpdateTime = data.UpdateTime.UnixMilli()
		temp.CreateTime = data.CreateTime.UnixMilli()
		respData = append(respData, temp)
	}

	return &pb.SearchResCommentByUserOrResIdResp{
		ResComment: respData,
	}, nil
}
