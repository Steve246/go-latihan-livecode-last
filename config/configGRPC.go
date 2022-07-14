package config

import "os"

type GrpcConfig struct {
	Url string
}
type ConfigGrpc struct {
	GrpcConfig
}

func (c *ConfigGrpc) readConfig() {
	grpcUrl := os.Getenv("GRPC_URL") //set GRPC_URL=localhost:8888
	c.GrpcConfig = GrpcConfig{Url: grpcUrl}
}
func NewConfigGrpc() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
