package pcfmicro

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-web"
	"github.com/st3v/cfkit/env"
	"github.com/st3v/pcf-micro/springcloud"
)

func init() {
	registry.DefaultRegistry = springcloud.NewRegistry()
	server.DefaultServer = NewServer()
	client.DefaultClient = NewClient()
	selector.DefaultSelector = NewSelector()
}

func NewWebService(opts ...web.Option) web.Service {
	opts = append([]web.Option{
		web.Address(env.Addr()),
		web.Advertise(advertise()),
	}, opts...)

	return web.NewService(opts...)
}

func advertise() string {
	app, err := env.Application()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s:80", app.URIs[0])
}

func NewService(opts ...micro.Option) micro.Service {
	opts = append([]micro.Option{
		micro.RegisterInterval(10 * time.Second),
		micro.Server(NewServer()),
		micro.Client(NewClient()),
		micro.Registry(springcloud.NewRegistry()),
	}, opts...)

	return micro.NewService(opts...)
}

func NewServer(opts ...server.Option) server.Server {
	opts = append([]server.Option{
		server.Address(env.Addr()),
		server.Advertise(advertise()),
		server.Registry(springcloud.NewRegistry()),
	}, opts...)

	return server.NewServer(opts...)
}

func NewClient(opts ...client.Option) client.Client {
	opts = append([]client.Option{
		client.Registry(springcloud.NewRegistry()),
		client.Selector(NewSelector()),
	}, opts...)

	return client.NewClient(opts...)
}

func NewSelector(opts ...selector.Option) selector.Selector {
	opts = append([]selector.Option{
		selector.Registry(springcloud.NewRegistry()),
	}, opts...)

	return selector.NewSelector(opts...)
}
