// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"example/internal/logic"
	"example/internal/svc"
	"example/types/pb"
)

type OrderServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServiceServer
}

func NewOrderServiceServer(svcCtx *svc.ServiceContext) *OrderServiceServer {
	return &OrderServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServiceServer) GetUserInfo(ctx context.Context, in *pb.UserReq) (*pb.UserResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
