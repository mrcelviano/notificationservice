package grpc

import (
	"context"
	"github.com/mrcelviano/notificationservice/internal/app"
	p "github.com/mrcelviano/notificationservice/proto"
	"github.com/pkg/errors"
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
	taskID, err := g.logic.RegisterTask(ctx, app.Task{
		Email: req.User.Email,
		Name:  req.User.Name,
	})
	if err != nil {
		return &p.SendNotificationResponse{TaskID: 0}, errors.Wrap(err, "can`t send notification")
	}
	return &p.SendNotificationResponse{TaskID: taskID}, nil
}
