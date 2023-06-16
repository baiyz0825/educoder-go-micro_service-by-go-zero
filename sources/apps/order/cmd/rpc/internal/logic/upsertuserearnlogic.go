package logic

import (
	"context"

	"gorm.io/gorm/clause"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpsertUserEarnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpsertUserEarnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpsertUserEarnLogic {
	return &UpsertUserEarnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpsertUserEarn
//
//	@Description: 插入或者更新userEarn
//	@receiver l
//	@param in
//	@return *pb.AddUserEarnResp
//	@return error
func (l *UpsertUserEarnLogic) UpsertUserEarn(in *pb.AddUserEarnReq) (*pb.AddUserEarnResp, error) {
	earn := l.svcCtx.Query.UserEarn
	userEarn := &model.UserEarn{
		UserID:  in.UserId,
		EarnNum: in.EarnNum,
		PayNum:  in.PayNum,
	}
	var conflict clause.OnConflict
	if in.EarnNum != 0 && in.PayNum != 0 {
		conflict = clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"pay_num", "earn_num"}),
		}
	} else if in.EarnNum != 0 {
		conflict = clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"earn_num"}),
		}
	} else if in.PayNum != 0 {
		conflict = clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"pay_num"}),
		}
	} else {
		return &pb.AddUserEarnResp{}, nil
	}
	err := earn.WithContext(context.Background()).Clauses(conflict).Create(userEarn)
	if err != nil {
		l.Logger.WithFields(logx.Field("userEarn:", userEarn)).Error("插入更新失败")
	}
	return &pb.AddUserEarnResp{}, nil
}
