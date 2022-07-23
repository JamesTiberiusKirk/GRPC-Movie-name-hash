package main

import (
	"log"
	"net"

	"github.com/JamesTiberiusKirk/moviehash/common/hashmovieservice"
	"github.com/JamesTiberiusKirk/moviehash/cryptoservice/controllers"
	"google.golang.org/grpc"
)

func runServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	log.Printf("Listening on %s", lis.Addr())

	hashMovieNameServer := controllers.NewHashMovieNameServer()
	hashmovieservice.RegisterHashMovieNameServer(server, hashMovieNameServer)

	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
