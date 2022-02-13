package delivery

import (
	"context"
	"github.com/mrcelviano/notificationservice/internal/domain"
	server "github.com/mrcelviano/notificationservice/pkg/notification/proto"
	"google.golang.org/grpc"
)

type notificationServer struct {
	service domain.NotificationService
}

func NewNotificationServer(service domain.NotificationService, opts ...grpc.ServerOption) *grpc.Server {
	g := &notificationServer{service: service}

	grpcServer := grpc.NewServer(opts...)
	server.RegisterNotificationServiceServer(grpcServer, g)
	return grpcServer
}

func (g *notificationServer) RegisterNotification(ctx context.Context, req *server.RegisterNotificationRequest) (resp *server.RegisterNotificationResponse, err error) {
	isRegistered, err := g.service.RegisterTask(ctx, req.UserID)
	if err != nil {
		return &server.RegisterNotificationResponse{}, domain.ErrRegisterNotification
	}
	return &server.RegisterNotificationResponse{IsRegistered: isRegistered}, nil
}
