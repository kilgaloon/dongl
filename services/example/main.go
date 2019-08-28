package example

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kilgaloon/dongl/agent"
	"github.com/kilgaloon/dongl/api"
	"github.com/kilgaloon/dongl/daemon"

	"github.com/kilgaloon/dongl/config"
)

// Client settings and configurations
type Client struct {
	Name string
	*agent.Default
}

// New create client as a service
func (client *Client) New(name string, cfg config.AgentConfig, debug bool) daemon.Service {
	a := agent.New(name, cfg, debug)
	c := &Client{
		name,
		a,
	}

	return c
}

// RName returns registered name for this service
func (client Client) RName() string {
	return client.Name
}

// Start client
func (client *Client) Start() {
	client.SetStatus("Started")
	fmt.Print("I Started!")
	go func() {
		for {
			if client.Status() == "Stopped" {
				fmt.Print("I Stopped!")
				break
			} else {
				fmt.Print("tick...")
			}

			time.Sleep(time.Second * 1)
		}
	}()
}

// Stop client
func (client *Client) Stop() {
	client.SetStatus("Stopped")
}

// RegisterAPIHandles to be used in http communication
func (client *Client) RegisterAPIHandles() map[string]api.Handle {
	cmds := make(map[string]api.Handle)

	cmds["stop"] = func(w http.ResponseWriter, r *http.Request) {
		client.Stop()
	}

	cmds["start"] = func(w http.ResponseWriter, r *http.Request) {
		client.Start()
	}

	return cmds
}
