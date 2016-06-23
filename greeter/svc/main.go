package main

import (
	"fmt"
	"log"

	"github.com/micro/go-micro"
	"golang.org/x/net/context"

	"github.com/st3v/cfkit/env"
	"github.com/st3v/pcf-micro"

	proto "github.com/st3v/pcf-micro-examples/greeter/proto"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Greeting = fmt.Sprintf("Hello from instance-%d, %s!", app.Instance.Index, req.Name)
	return nil
}

var app env.App

func main() {
	var err error
	app, err = env.Application()
	if err != nil {
		log.Fatal(err)
	}

	svc := pcfmicro.NewService(
		micro.Name("go.micro.svc.greeter"),
		micro.Version("0.0.1"),
	)

	svc.Init()

	proto.RegisterSayHandler(svc.Server(), new(Say))

	if err := svc.Run(); err != nil {
		log.Fatalf("Error running service: %s", err)
	}
}
