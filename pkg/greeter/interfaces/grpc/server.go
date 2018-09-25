package grpc

import (
	"context"
	"log"

	"github.com/charliekenney23/grpc-greet-service/pkg/greeter/infrastructure/proto"
)

type greetingServer struct{}

func (s *greetingServer) SayHello(ctx context.Context, in *proto.Greeting) (*proto.Greeting, error) {
	log.Printf("user said: '%s'", in.Contents)
	return &proto.Greeting{
		Contents: "Hello!",
	}, nil
}

// New returns a server instance as a new grpc
// gateway interface
func New() proto.GreeterServer {
	return &greetingServer{}
}
