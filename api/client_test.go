package api

import (
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type RegistratorMock struct {
}

func (rm RegistratorMock) RName() string {
	return "registrator_mock"
}

// RegisterAPIHandles to be used in socket communication
// If you want to takeover default commands from agent
// call DefaultCommands from Agent which is same command
func (rm *RegistratorMock) RegisterAPIHandles() map[string]Handle {
	cmds := make(map[string]Handle)

	cmds["hello"] = func(w http.ResponseWriter, r *http.Request) {
		var resp struct {
			Message string
		}

		resp.Message = "Hello world!"

		j, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)
	}

	cmds["info"] = func(w http.ResponseWriter, r *http.Request) {
		resp := InfoResponse{
			Status:         "testing",
			RecipesInQueue: "10",
		}

		w.WriteHeader(http.StatusOK)

		j, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)

		return
	}

	cmds["stop"] = func(w http.ResponseWriter, r *http.Request) {
		resp := MessageResponse{
			Message: "Mock stopped",
		}

		j, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)
	}

	return cmds
}

func (rm *RegistratorMock) DefaultAPIHandles() map[string]Handle {
	cmds := make(map[string]Handle)

	cmds["hello_def"] = func(w http.ResponseWriter, r *http.Request) {
		var resp struct {
			Message string
		}

		resp.Message = "Hello world DEFAULT!"

		j, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)
	}

	return cmds
}

var rm = &RegistratorMock{}

func TestMain(t *testing.T) {
	assert.Equal(t, IsAPIRunning(), false)
	a := New()

	a.RegisterHandle("hello_reg", func(w http.ResponseWriter, r *http.Request) {
		var resp struct {
			Message string
		}

		resp.Message = "Hello world REGISTERED!"

		j, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(j)
	})

	go a.Register(rm).Start()
	assert.Equal(t, IsAPIRunning(), true)
}
func TestCmd(t *testing.T) {
	c := Cmd("agent command arg")

	assert.Equal(t, c.Agent(), "agent")
	assert.Equal(t, c.Command(), "command")
	assert.Equal(t, c.Args()[0], "arg")
}

func TestResolver(t *testing.T) {
	Resolver(Cmd("registrator_mock not_exist"))
	Resolver(Cmd("registrator_mock"))
	Resolver(Cmd("registrator_mock info"))
	Resolver(Cmd("registrator_mock stop"))
	Resolver(Cmd("registrator_mock workers:list"))
	Resolver(Cmd("registrator_mock workers:kill job"))

	Resolver(Cmd("agent workers:kill job"))
	Resolver(Cmd("agent info"))
	Resolver(Cmd("agent stop"))
	Resolver(Cmd("agent workers:list"))
}

func TestRevealEndpoint(t *testing.T) {
	assert.Equal(t, "http://localhost:11401/agent/command", RevealEndpoint("/{agent}/command", Cmd("agent command")))
}
