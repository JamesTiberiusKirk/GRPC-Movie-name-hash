package main

import (
	"os"
)

type Config struct {
	GrpcPort string
}

const (
	grpcPortEnv = "GRPC_PORT"
)

func buildConfig() Config {
	return Config{
		GrpcPort: os.Getenv(grpcPortEnv),
	}
}
