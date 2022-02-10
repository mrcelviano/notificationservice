package commons

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

func StartGrpc(s *grpc.Server, port int) {
	log.Println("START GRPC SERVER ON PORT ", port)
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err, "can't listen tcp on port ", port)
	}
	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatal(err, "can't serve grpc server")
		}
	}()
}
