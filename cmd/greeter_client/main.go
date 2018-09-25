package main

import (
	"context"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	greeter_grpc "github.com/charliekenney23/grpc-greet-service/pkg/greeter/infrastructure/proto"
	"google.golang.org/grpc"
)

type config struct {
	GreeterHost string `env:"ENV"  envDefault:"0.0.0.0"`
	GreeterPort int    `env:"PORT" envDefault:"3000"`
}

func main() {
	cfg := &config{}
	env.Parse(cfg)

	greeterConn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.GreeterHost,
		cfg.GreeterPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal("error dialing greeter service", err.Error())
	}

	greeterClient := greeter_grpc.NewGreeterClient(greeterConn)
	ctx := context.Background()
	res, err := greeterClient.SayHello(ctx, &greeter_grpc.Greeting{Contents: "Bonjour!"})
	if err != nil {
		log.Fatal("error greeting the greeter service", err.Error())
	}

	log.Println("got greeting from greeter service =>", res.Contents)
}
