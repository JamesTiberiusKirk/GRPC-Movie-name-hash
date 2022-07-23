package main

import "log"

const (
	port = ":3000"
)

func main() {
	config := buildConfig()
	log.Printf("Config: %+v", config)

	err := runServer(config.GrpcPort)
	if err != nil {
		log.Fatalf("Error running the server %v", err)
	}
}
