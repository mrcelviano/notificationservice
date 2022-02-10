package grpc

import (
	"context"
	"fmt"
	"github.com/mrcelviano/notificationservice/app"
	p "github.com/mrcelviano/notificationservice/proto"
	"google.golang.org/grpc"
)

type grpcHandlers struct {
	logic app.NotificationLogic
}

func NewGRPCHandlers(logic app.NotificationLogic, opts ...grpc.ServerOption) *grpc.Server {
	g := &grpcHandlers{logic: logic}

	grpcServer := grpc.NewServer(opts...)
	p.RegisterNotificationServiceServer(grpcServer, g)
	return grpcServer
}

func (g *grpcHandlers) SendNotification(ctx context.Context, req *p.SendNotificationRequest) (resp *p.SendNotificationResponse, err error) {
	fmt.Println("New Request!!!")
	return &p.SendNotificationResponse{TaskID: 0}, nil
}
