package main

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/micro/go-micro/metadata"
	"github.com/st3v/pcf-micro"

	proto "github.com/st3v/pcf-micro-examples/greeter/proto"
)

func main() {
	// Create the client
	client := pcfmicro.NewClient()

	// Create new request to service go.micro.srv.greeter, method Say.Hello
	req := client.NewRequest("go.micro.svc.greeter", "Say.Hello", &proto.Request{
		Name: "John",
	})

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &proto.Response{}

	// Call service
	if err := client.Call(ctx, req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Greeting)
}
