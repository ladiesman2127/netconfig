package client

import (
	"log"

	pb "github.com/ladiesman2127/netconfig/api/gen/go"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() (pb.NetConfigClient, *grpc.ClientConn) {
	connection, err := grpc.NewClient(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Unable to connect to gRPC server: %s", err.Error())
	}
	return pb.NewNetConfigClient(connection), connection
}
