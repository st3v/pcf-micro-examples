package springcloud

import (
	"log"
	"strings"

	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/eureka"
	"github.com/st3v/cfkit/env"
)

func NewRegistry(opts ...registry.Option) registry.Registry {
	service, err := env.ServiceWithTag("eureka")
	if err != nil {
		log.Println("Service with tag 'registry' not found.")
		return nil
	}

	registryURL, ok := service.Credentials["uri"].(string)
	if !ok {
		log.Printf("Invalid registry URI '%v'", registryURL)
		return nil
	}

	if !strings.HasSuffix(registryURL, "/eureka") {
		registryURL = strings.Join([]string{registryURL, "/eureka"}, "")
	}

	clientID, ok := service.Credentials["client_id"].(string)
	if !ok {
		log.Printf("Invalid OAuth2 ClientID '%v'", clientID)
		return nil
	}

	clientSecret, ok := service.Credentials["client_secret"].(string)
	if !ok {
		log.Printf("Invalid OAuth2 ClientSecret '%v'", clientSecret)
		return nil
	}

	tokenURL, ok := service.Credentials["access_token_uri"].(string)
	if !ok {
		log.Printf("Invalid OAuth2 TokenURL '%v'", tokenURL)
		return nil
	}

	opts = append([]registry.Option{
		registry.Addrs(registryURL),
		eureka.OAuth2ClientCredentials(clientID, clientSecret, tokenURL),
	}, opts...)

	return eureka.NewRegistry(opts...)
}
