package main

import (
	"fmt"
	"log"
	"net"

	"github.com/caarlos0/env"
	"github.com/charliekenney23/grpc-greet-service/pkg/greeter/infrastructure/proto"
	server "github.com/charliekenney23/grpc-greet-service/pkg/greeter/interfaces/grpc"
	"google.golang.org/grpc"
)

type config struct {
	Env  string `env:"ENV"  envDefault:"development"`
	Host string `env:"HOST"  envDefault:"0.0.0.0"`
	Port int    `env:"PORT" envDefault:"3000"`
}

func main() {
	cfg := &config{}
	env.Parse(cfg)

	grpcSrv := grpc.NewServer()
	greeter := server.New()
	proto.RegisterGreeterServer(grpcSrv, greeter)

	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal("failed to start tcp server")
	}

	grpcSrv.Serve(ln)
}
