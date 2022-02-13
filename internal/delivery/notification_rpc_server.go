package delivery

import (
	"context"
	"github.com/mrcelviano/notificationservice/internal/domain"
	"github.com/mrcelviano/notificationservice/pkg/notification/proto"
	"google.golang.org/grpc"
)

type notificationServer struct {
	service domain.NotificationService
}

func NewNotificationServer(service domain.NotificationService, opts ...grpc.ServerOption) *grpc.Server {
	g := &notificationServer{service: service}

	grpcServer := grpc.NewServer(opts...)
	proto.RegisterNotificationServiceServer(grpcServer, g)
	return grpcServer
}

func (g *notificationServer) SendNotification(ctx context.Context, req *proto.SendNotificationRequest) (resp *proto.SendNotificationResponse, err error) {
	taskID, err := g.service.RegisterTask(ctx, domain.Task{
		Email: req.User.Email,
		Name:  req.User.Name,
	})
	if err != nil {
		return &proto.SendNotificationResponse{TaskID: 0}, domain.ErrRegisterNotification
	}
	return &proto.SendNotificationResponse{TaskID: taskID}, nil
}
