package agent

import (
	"io"
	"os"
	"sync"

	"github.com/kilgaloon/dongl/api"
	"github.com/kilgaloon/dongl/config"
)

// Agent interface defines service that can be started/stop
// that has workers, config, context
type Agent interface {
	Config() *config.AgentConfig
	Status() string
	SetStatus(s string)
}

// Default represents default agent
type Default struct {
	Name   string
	config config.AgentConfig
	Stdin  io.Reader
	Stdout io.Writer
	Debug  bool
	status string

	*sync.RWMutex
}

// RName returns agent name
func (d Default) RName() string {
	return d.Name
}

// Config return current config for agent
func (d Default) Config() config.AgentConfig {
	return d.config
}

func (d Default) Write(p []byte) (n int, err error) {
	return d.Stdout.Write(p)
}

func (d Default) Read(p []byte) (n int, err error) {
	return d.Stdin.Read(p)
}

// IsDebug determines is agent in debug mode
func (d Default) IsDebug() bool {
	return d.Debug
}

// SetStatus set status of agent
func (d *Default) SetStatus(s string) {
	d.status = s
}

// Status return current status of agent
func (d *Default) Status() string {
	return d.status
}

// DefaultAPIHandles to be used in socket communication
// If you want to takeover default commands from agent
// call DefaultCommands from Agent which is same command
func (d *Default) DefaultAPIHandles() map[string]api.Handle {
	cmds := make(map[string]api.Handle)

	// this function merge both maps and inject default commands from agent
	return cmds
}

// New default client
func New(name string, cfg config.AgentConfig, debug bool) *Default {
	agent := &Default{}
	agent.config = cfg
	agent.RWMutex = new(sync.RWMutex)
	agent.Stdin = os.Stdin
	agent.Stdout = os.Stdout
	agent.Debug = debug

	return agent
}
