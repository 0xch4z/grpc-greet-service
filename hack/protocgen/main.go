package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type service struct {
	name      string
	protoPath string
}

// todo(charles): use protolib programatically rather than
// running extern commands via shell
func (s *service) GenerateProto() {
	cmd := exec.Command("protoc", "--proto_path="+s.protoPath,
		"--go_out=plugins=grpc:"+s.protoPath, s.name+".proto")
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fatal("failed to generate proto for service:", s.name, err.Error())
	}
}

func newProtoService(name string) *service {
	path, err := filepath.Abs("pkg/" + name + "/infrastructure/proto")
	if err != nil {
		fatal("failed to make proto path", err.Error())
	}
	return &service{name, path}
}

var services = []*service{
	newProtoService("greeter"),
}

func fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format, args...)
	os.Exit(1)
}

func main() {
	wg := sync.WaitGroup{}
	for _, svc := range services {
		wg.Add(1)
		go func(s *service) {
			defer wg.Done()
			s.GenerateProto()
		}(svc)
	}

	wg.Wait()
}
