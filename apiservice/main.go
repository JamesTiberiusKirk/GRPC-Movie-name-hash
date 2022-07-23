package main

import (
	"log"
)

const (
	address = "localhost:3000"
)

func main() {
	config := buildConfig()
	log.Printf("Config: %+v", config)

	grpcClinet, conn, err := initGrpcClient(config.GrpcAddress, config.GrpcPort)
	if err != nil {
		log.Fatalf("Error initialising grpcClinet %v", err)
	}
	defer conn.Close()

	runServer(config.HttpPort, grpcClinet)
}
