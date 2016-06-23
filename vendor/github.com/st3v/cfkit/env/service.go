package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const svcEnvVar = "VCAP_SERVICES"

type Service struct {
	Name        string                 `json:"name"`
	Label       string                 `json:"label"`
	Tags        []string               `json:"tags"`
	Plan        string                 `json:"plan"`
	Credentials map[string]interface{} `json:"credentials"`
}

func ServiceWithTag(tag string) (Service, error) {
	services, err := getServices()
	if err != nil {
		return Service{}, err
	}

	return services.withTag(tag)
}

func ServiceWithName(name string) (Service, error) {
	services, err := getServices()
	if err != nil {
		return Service{}, err
	}

	return services.withName(name)
}

type serviceMap map[string][]Service

func getServices() (serviceMap, error) {
	jsonStr := os.Getenv(svcEnvVar)
	if jsonStr == "" {
		return serviceMap{}, fmt.Errorf("%s not set", svcEnvVar)
	}

	services := new(serviceMap)
	if err := json.Unmarshal([]byte(jsonStr), services); err != nil {
		return serviceMap{}, fmt.Errorf("Error parsing %s: %s", svcEnvVar, err)
	}

	return *services, nil
}

func (m serviceMap) withTag(tag string) (Service, error) {
	for _, services := range m {
		for _, service := range services {
			for _, t := range service.Tags {
				if strings.ToUpper(t) == strings.ToUpper(tag) {
					return service, nil
				}
			}
		}
	}
	return Service{}, fmt.Errorf("Service with tag '%s' not found", tag)
}

func (m serviceMap) withName(name string) (Service, error) {
	for _, services := range m {
		for _, service := range services {
			if strings.ToUpper(name) == strings.ToUpper(service.Name) {
				return service, nil
			}
		}
	}
	return Service{}, fmt.Errorf("Service with name '%s' not found", name)
}
