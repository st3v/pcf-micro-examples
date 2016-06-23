package main

import (
	"encoding/json"
	"log"
	"strings"

	"golang.org/x/net/context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/server"
	api "github.com/micro/micro/api/proto"

	_ "github.com/st3v/pcf-micro"

	proto "github.com/st3v/pcf-micro-examples/greeter/proto"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Say.Hello API request")

	name, ok := req.Get["name"]
	if !ok || len(name.Values) == 0 {
		return errors.BadRequest("go.micro.api.greeter", "Name cannot be blank")
	}

	request := client.NewRequest("go.micro.svc.greeter", "Say.Hello", &proto.Request{
		Name: strings.Join(name.Values, " "),
	})

	response := &proto.Response{}

	if err := client.Call(ctx, request, response); err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"message": response.Greeting,
	})
	rsp.Body = string(b)

	return nil
}

func main() {
	// Set server name
	server.Init(
		server.Name("go.micro.api.greeter"),
	)

	// Register Handlers
	server.Handle(
		server.NewHandler(
			new(Say),
		),
	)

	// Run server
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
