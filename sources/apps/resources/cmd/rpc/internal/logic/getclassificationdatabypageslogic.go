package logic

import (
	"context"
	"sync"

	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/svc"
	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/pb"
	"github.com/baiyz0825/school-share-buy-backend/common/xconst"
	"github.com/baiyz0825/school-share-buy-backend/common/xerr"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassificationDataByPagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetClassificationDataByPagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassificationDataByPagesLogic {
	return &GetClassificationDataByPagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetClassificationDataByPages 分类查询分类id下的资源
func (l *GetClassificationDataByPagesLogic) GetClassificationDataByPages(in *pb.ClassificationDataByPagesReq) (*pb.ClassificationDataByPagesResp, error) {
	// check pb
	if in == nil {
		return nil, xerr.NewErrCode(xerr.PB_CHECK_ERR)
	}
	// if in.GetClassificationID() == 0 {
	// 	return nil, xerr.NewErrCode(xerr.PB_LOGIC_CHECK_ERR)
	// }
	// pages 判断
	if in.GetPage() <= 0 || in.GetLimit() < 0 {
		return nil, xerr.NewErrCode(xerr.RPC_PAGES_PARAM_ERR)
	}
	// search db
	// files
	file := l.svcCtx.Query.File
	// texts
	text := l.svcCtx.Query.OnlineText
	var fileFindsData []*model.File
	var textFindsData []*model.OnlineText
	var errFile, errText error
	resp := &pb.ClassificationDataByPagesResp{}
	if in.GetResType() == xconst.TEXT_TYPE {
		var condition []gen.Condition
		// 全查所有分类数据
		if in.GetClassificationID() != 0 {
			condition = append(condition, text.ClassID.Eq(in.GetClassificationID()))
		}
		// 文本类型
		if in.GetUserId() != 0 {
			condition = append(condition, text.Owner.Eq(in.GetUserId()))
		} else {
			condition = append(condition, text.Permission.Eq(xconst.PERMISSION_TRUE))
		}
		// 是否存在关键词
		if len(in.KeyWord) > 0 {
			condition = append(condition, text.TextName.Like(in.KeyWord))
		}
		textFindsData, total, errText := text.WithContext(l.ctx).Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
		if errText != nil && errText != gorm.ErrRecordNotFound {
			l.Logger.WithFields(
				logx.LogField{
					Key:   "fileErr",
					Value: errFile,
				},
				logx.LogField{
					Key:   "textErr",
					Value: errText,
				}).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		var respText []*pb.OnlineText
		if textFindsData != nil {
			for _, data := range textFindsData {
				temp := &pb.OnlineText{}
				err := copier.Copy(temp, data)
				if err != nil {
					return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
				}
				temp.UpdateTime = data.UpdateTime.UnixMilli()
				temp.CreateTime = data.CreateTime.UnixMilli()
				respText = append(respText, temp)
			}
			resp.OnlineText = respText
			// 总数
			resp.TextsTotal = total
		} else {
			resp.OnlineText = nil
		}
	} else if in.GetResType() == xconst.FILE_TYEP {
		// 文件类型
		var condition []gen.Condition
		if in.GetClassificationID() != 0 {
			condition = append(condition, file.Class.Eq(in.GetClassificationID()))
		}
		if in.GetUserId() != 0 {
			condition = append(condition, file.Owner.Eq(in.GetUserId()))
		} else {
			condition = append(condition, file.DownloadAllow.Eq(xconst.PERMISSION_TRUE))
		}
		// 关键词
		if len(in.KeyWord) > 0 {
			condition = append(condition, file.Name.Like(in.KeyWord))
		}
		fileFindsData, total, errFile := file.WithContext(l.ctx).Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
		if errFile != nil && errFile != gorm.ErrRecordNotFound {
			l.Logger.WithFields(
				logx.LogField{
					Key:   "fileErr",
					Value: errFile,
				},
				logx.LogField{
					Key:   "textErr",
					Value: errText,
				}).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}
		if fileFindsData != nil {
			var respFiles []*pb.File
			// copy data
			for _, data := range fileFindsData {
				temp := &pb.File{}
				err := copier.Copy(temp, data)
				if err != nil {
					return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
				}
				temp.UpdateTime = data.UpdateTime.UnixMilli()
				temp.CreateTime = data.CreateTime.UnixMilli()
				respFiles = append(respFiles, temp)
			}
			resp.Files = respFiles
			resp.FilesTotal = total
		} else {
			resp.Files = nil
		}
	} else {
		// 获取全部
		var wg sync.WaitGroup
		var totalFiles int64 = 0
		var totalText int64 = 0
		wg.Add(2)
		go func() {
			var condition []gen.Condition
			if in.GetClassificationID() != 0 {
				condition = append(condition, file.Class.Eq(in.GetClassificationID()))
			}
			if in.GetUserId() != 0 {
				condition = append(condition, file.Owner.Eq(in.GetUserId()))
			} else {
				condition = append(condition, file.DownloadAllow.Eq(xconst.PERMISSION_TRUE))
			}
			// 关键词
			if len(in.KeyWord) > 0 {
				condition = append(condition, file.Name.Like(in.KeyWord))
			}
			data, total, errFileTemp := file.WithContext(l.ctx).Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
			fileFindsData = data
			errFile = errFileTemp
			totalFiles = total
			wg.Done()
		}()
		go func() {
			var condition []gen.Condition
			// 全查所有分类数据
			if in.GetClassificationID() != 0 {
				condition = append(condition, text.ClassID.Eq(in.GetClassificationID()))
			}
			if in.GetUserId() != 0 {
				condition = append(condition, text.Owner.Eq(in.GetUserId()))
			} else {
				condition = append(condition, text.Permission.Eq(xconst.PERMISSION_TRUE))
			}
			// 是否存在关键词
			if len(in.KeyWord) > 0 {
				condition = append(condition, text.TextName.Like(in.KeyWord))
			}
			data, total, errTextTemp := text.WithContext(l.ctx).Where(condition...).FindByPage(int((in.GetPage()-1)*in.GetLimit()), int(in.GetLimit()))
			textFindsData = data
			totalText = total
			errText = errTextTemp
			wg.Done()
		}()
		wg.Wait()

		if (errFile != nil && errFile != gorm.ErrRecordNotFound) || (errText != nil && errText != gorm.ErrRecordNotFound) {
			l.Logger.WithFields(
				logx.LogField{
					Key:   "fileErr",
					Value: errFile,
				},
				logx.LogField{
					Key:   "textErr",
					Value: errText,
				}).Error(xerr.NewErrCode(xerr.RPC_SEARCH_ERR))
			return nil, xerr.NewErrCode(xerr.RPC_SEARCH_ERR)
		}

		// return var
		var respFiles []*pb.File
		var respText []*pb.OnlineText
		if fileFindsData != nil {
			// copy data
			for _, data := range fileFindsData {
				temp := &pb.File{}
				err := copier.Copy(temp, data)
				if err != nil {
					return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
				}
				temp.UpdateTime = data.UpdateTime.UnixMilli()
				temp.CreateTime = data.CreateTime.UnixMilli()
				respFiles = append(respFiles, temp)
			}
			resp.Files = respFiles
			resp.FilesTotal = totalFiles
		} else {
			resp.Files = nil
		}

		if textFindsData != nil {
			for _, data := range textFindsData {
				temp := &pb.OnlineText{}
				err := copier.Copy(temp, data)
				if err != nil {
					return nil, errors.Wrapf(err, xerr.GetErrMsg(xerr.RPC_SEARCH_ERR))
				}
				temp.UpdateTime = data.UpdateTime.UnixMilli()
				temp.CreateTime = data.CreateTime.UnixMilli()
				respText = append(respText, temp)
			}
			resp.OnlineText = respText
			resp.TextsTotal = totalText
		} else {
			resp.OnlineText = nil
		}
	}
	// return
	return resp, nil
}
