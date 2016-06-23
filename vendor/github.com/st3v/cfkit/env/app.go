package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const appEnvVar = "VCAP_APPLICATION"

type App struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	URIs           []string    `json:"uris"`
	Version        string      `json:"version"`
	Host           string      `json:"host"`
	Port           int         `json:"port"`
	Addr           string      `json:"addr"`
	Limits         AppLimits   `json:"limits"`
	StartTimestamp int         `json:"started_at_timestamp"`
	StateTimestamp int         `json:"state_timestamp"`
	Instance       AppInstance `json:"instance"`
	Space          Space       `json:"space"`
}

type AppLimits struct {
	Disk            int `json:"disk"`
	Memory          int `json:"mem"`
	FileDescriptors int `json:"fds"`
}

type AppInstance struct {
	ID    string `json:"id"`
	Index int    `json:"index"`
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Addr  string `json:"addr"`
}

type Space struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Application() (App, error) {
	vcapApp := os.Getenv(appEnvVar)
	if vcapApp == "" {
		return App{}, fmt.Errorf("%s not set", appEnvVar)
	}

	var app App
	if err := json.Unmarshal([]byte(vcapApp), &app); err != nil {
		return App{}, fmt.Errorf("Error parsing %s: %s", appEnvVar, err)
	}

	return app, nil
}

func (a *App) URI() string {
	if len(a.URIs) == 0 {
		return ""
	}
	return a.URIs[0]
}

func (a *App) UnmarshalJSON(data []byte) error {
	type AppAlias App

	aux := &struct {
		InstanceID    string `json:"instance_id"`
		InstanceIndex int    `json:"instance_index"`
		AppID         string `json:"application_id"`
		SpaceID       string `json:"space_id"`
		SpaceName     string `json:"space_name"`
		*AppAlias
	}{
		AppAlias: (*AppAlias)(a),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	a.ID = aux.AppID
	a.Space.ID = aux.SpaceID
	a.Space.Name = aux.SpaceName
	a.Addr = fmt.Sprintf("%s:%d", aux.Host, aux.Port)
	a.Instance = AppInstance{
		ID:    aux.InstanceID,
		Index: aux.InstanceIndex,
		Port:  instancePort(),
		IP:    instanceIP(),
		Addr:  instanceAddr(),
	}

	return nil
}

func instancePort() int {
	port, err := strconv.Atoi(os.Getenv("CF_INSTANCE_PORT"))
	if err != nil {
		return 0
	}
	return port
}

func instanceIP() string {
	return os.Getenv("CF_INSTANCE_IP")
}

func instanceAddr() string {
	return os.Getenv("CF_INSTANCE_ADDR")
}
