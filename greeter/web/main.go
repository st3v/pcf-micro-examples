package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-web"

	"github.com/st3v/pcf-micro"

	proto "github.com/st3v/pcf-micro-examples/greeter/proto"
)

func main() {
	service := pcfmicro.NewWebService(
		web.Name("go.micro.web.greeter"),
	)

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.Form.Get("name")
			if len(name) == 0 {
				name = "World"
			}

			cl := proto.NewSayClient("go.micro.svc.greeter", client.DefaultClient)
			rsp, err := cl.Hello(context.Background(), &proto.Request{
				Name: name,
			})

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}

			w.Write([]byte(`<html><body><h1>` + rsp.Greeting + `</h1></body></html>`))
			return
		}

		fmt.Fprint(w, `<html><body><h1>Enter Name<h1><form method=post><input name=name type=text /></form></body></html>`)
	})

	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
