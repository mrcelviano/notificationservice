package grpc

import (
	"google.golang.org/grpc"
	"social-tech/notificationservice/app"
)

type grpcHandlers struct {
	logic app.NotificationLogic
}

func NewGRPCHandlers(logic app.NotificationLogic, opts ...grpc.ServerOption) *grpc.Server {
	return nil
}
