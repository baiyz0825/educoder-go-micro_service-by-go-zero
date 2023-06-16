package bo

// OrderMqStruct OrderMqStruct
//
//	@Description: 订单定时任务描述
type OrderMqStruct struct {
	Uuid            int64
	UserId          int64
	PayPathOrderNum string
}
