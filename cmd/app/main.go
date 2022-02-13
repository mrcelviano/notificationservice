package main

import (
	"fmt"
	"github.com/mrcelviano/notificationservice/internal/config"
	"github.com/mrcelviano/notificationservice/internal/delivery"
	"github.com/mrcelviano/notificationservice/internal/repository"
	"github.com/mrcelviano/notificationservice/internal/service"
	"github.com/mrcelviano/notificationservice/pkg/database/postgres"
	"github.com/mrcelviano/notificationservice/pkg/logger"
	"github.com/mrcelviano/notificationservice/pkg/user"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	configsDirectory = "configs"
)

func main() {
	cfg, err := config.Init(configsDirectory)
	if err != nil {
		logger.Error(err)
		return
	}

	postgresConnection, err := postgres.NewGoCraftDBConnectionPG(cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User,
		cfg.Postgres.Password, cfg.Postgres.DBName)
	if err != nil {
		logger.Error(err)
		return
	}

	//repo
	var (
		notificationRepo = repository.NewNotificationRepositoryPG(postgresConnection)
	)

	//rpc
	userClient, err := user.NewUserClient(cfg.Services)
	if err != nil {
		logger.Error(err)
		return
	}

	//service
	var (
		senderService       = service.NewSenderService()
		notificationService = service.NewNotificationService(notificationRepo, senderService, userClient)
	)

	//delivery
	server := delivery.NewNotificationServer(notificationService)

	logger.Info("server start")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPC.Port))
	if err != nil {
		logger.Errorf("can't listen tcp on port %s: %s\n", cfg.GRPC.Port, err.Error())
	}
	go func() {
		err := server.Serve(lis)
		if err != nil {
			logger.Errorf("can`t run grpc server: %s\n", err.Error())
			return
		}
	}()

	logger.Info("scheduler start")

	notificationService.StartNotificationScheduler()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	server.GracefulStop()
}
