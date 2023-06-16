package aliCallback

import (
	"net/http"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/logic/aliCallback"
	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/api/internal/svc"
)

func AlipayNoticeCallBackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := aliCallback.NewAlipayNoticeCallBackLogic(r.Context(), svcCtx)
		notification, err := svcCtx.AliPayClient.GetTradeNotification(r)
		if err != nil {
			svcCtx.AliPayClient.ACKNotification(w)
		}
		l.Notification = notification
		err = l.AlipayNoticeCallBack()
		// 支付宝默认返回
		if err != nil {
			svcCtx.AliPayClient.ACKNotification(w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("处理失败"))
		}
	}
}
