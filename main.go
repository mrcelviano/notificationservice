package main

import (
	"github.com/mrcelviano/notificationservice/commons"
	gRPC "github.com/mrcelviano/notificationservice/internal/grpc"
	"github.com/mrcelviano/notificationservice/internal/logic"
	"github.com/mrcelviano/notificationservice/internal/repository"
	"google.golang.org/grpc"
)

func main() {
	env := commons.GetEnvVar()
	commons.ConfigInit("configs/" + env + "_setting.json")
	pg := commons.InitGocraftDBRConnectionPG()

	//repo
	var (
		notificationRepo = repository.NewNotificationRepository(pg)
	)

	//logiс
	var (
		sendLogic         = logic.NewSendNotificationLogic()
		notificationLogic = logic.NewNotificationLogic(notificationRepo, sendLogic)
	)

	//grpc
	g := gRPC.NewGRPCHandlers(notificationLogic, grpc.ChainUnaryInterceptor(
		commons.GRPCDBRSessionPG(pg),
	))

	commons.NewSignalHandler(g)
	commons.StartGrpc(g, 8081)
	notificationLogic.Start()
	select {}
}
