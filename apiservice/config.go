package main

import "os"

type Config struct {
	HttpPort    string
	GrpcAddress string
	GrpcPort    string
}

const (
	httpPortEnv    = "HTTP_PORT"
	grpcAddressEnv = "GRPC_ADDRESS"
	grpcPortEnv    = "GRPC_PORT"
)

func buildConfig() Config {
	return Config{
		HttpPort:    os.Getenv(httpPortEnv),
		GrpcAddress: os.Getenv(grpcAddressEnv),
		GrpcPort:    os.Getenv(grpcPortEnv),
	}
}
