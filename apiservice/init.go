package main

import (
	"log"

	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"google.golang.org/grpc"
)

func initGrpcClient(address, port string) (hashmovieservice.HashMovieNameClient, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(address+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
		return nil, nil, err
	}

	c := hashmovieservice.NewHashMovieNameClient(conn)

	return c, conn, nil
}
