package main

import (
	"log"
	"net"

	pb "github.com/ladiesman2127/netconfig/api/gen/go"
	"github.com/ladiesman2127/netconfig/internal/service"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %s", err.Error())
	}
	s := service.New()

	grpcServer := grpc.NewServer()

	pb.RegisterNetConfigServer(grpcServer, s)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to server gRPC server on port 9000: %s", err.Error())
	}
}
