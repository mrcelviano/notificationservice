package user

import (
	"context"
	client "github.com/mrcelviano/notificationservice/pkg/user/proto"
	pool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

const (
	userServiceName = "userservice"
)

type userGRPC struct {
	pool *pool.Pool
}

func NewUserClient(serviceAddress map[string]string) (Service, error) {
	user := userGRPC{}
	userAddress, err := user.getServiceName(serviceAddress)
	if err != nil {
		return nil, err
	}
	err = user.newGRPCConnectionPool(userAddress, 10)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userGRPC) GetUserByID(ctx context.Context, userID int64) (User, error) {
	ctx, cansel := context.WithTimeout(ctx, time.Second*10)
	defer cansel()

	clientConn, err := u.pool.Get(ctx)
	if err != nil {
		return User{}, err
	}

	user, err := client.NewUserServiceClient(clientConn).GetUserByID(ctx,
		&client.GetUserByIDRequest{
			UserID: userID,
		})
	if err != nil {
		return User{}, ErrUserNotFound
	}

	return User{
		ID:    user.GetID(),
		Email: user.GetEmail(),
		Name:  user.GetName(),
	}, nil
}

func (u *userGRPC) SetIsRegisteredStatus(ctx context.Context, userID int64) (bool, error) {
	ctx, cansel := context.WithTimeout(ctx, time.Second*10)
	defer cansel()

	clientConn, err := u.pool.Get(ctx)
	if err != nil {
		return false, err
	}

	isSetStatus, err := client.NewUserServiceClient(clientConn).SetIsRegisteredStatus(ctx,
		&client.SetIsRegisteredStatusRequest{
			UserID: userID,
		})
	if err != nil {
		return false, ErrNotSetRegisterStatus
	}

	return isSetStatus.GetStatusIsSet(), nil
}

func (u *userGRPC) newGRPCConnectionPool(serviceAddress string, capacity int) (err error) {
	factory := func() (*grpc.ClientConn, error) {
		conn, err := grpc.Dial(serviceAddress,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithDefaultCallOptions(
				grpc.WaitForReady(true),
			),
		)
		if err != nil {
			return nil, err
		}
		return conn, err
	}
	u.pool, err = pool.New(factory, 0, capacity*10, time.Minute, time.Minute)
	if err != nil {
		return err
	}
	return nil
}

func (u *userGRPC) getServiceName(serviceAddress map[string]string) (string, error) {
	notificationAddress, ok := serviceAddress[userServiceName]
	if !ok {
		return "", ErrUserAddressNotFound
	}
	return notificationAddress, nil
}
