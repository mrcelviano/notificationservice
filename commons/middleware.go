package commons

import (
	"context"
	"github.com/gocraft/dbr"
	"google.golang.org/grpc"
)

func GRPCDBRSessionPG(db *dbr.Connection) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(DBSessionNewContext(ctx, db), req)
	}
}
