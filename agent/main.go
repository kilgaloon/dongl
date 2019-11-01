package agent

import (
	"io"
	"os"
	"sync"

	"github.com/kilgaloon/dongl/api"
	"github.com/spf13/viper"
)

// Agent interface defines service that can be started/stop
// that has workers, config, context
type Agent interface {
	Config() *viper.Viper
	Status() string
	SetStatus(s string)
}

// Default represents default agent
type Default struct {
	Name    string
	config  *viper.Viper
	Stdin   io.Reader
	Stdout  io.Writer
	Debug   bool
	status  string
	plugins map[string]Plugin

	*sync.RWMutex
}

// RName returns agent name
func (d Default) RName() string {
	return d.Name
}

// Config return current config for agent
func (d Default) Config() *viper.Viper {
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

// RegisterPlugin registers plugin and assing it to default agent
func (d *Default) RegisterPlugin(p Plugin) {
	d.plugins[p.Name()] = p.Bootstrap()
}

// Plugin returns registered plugin
func (d *Default) Plugin(n string) Plugin {
	return d.plugins[n]
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
func New(name string, cfg *viper.Viper, debug bool) *Default {
	agent := &Default{}
	agent.config = cfg
	agent.RWMutex = new(sync.RWMutex)
	agent.Stdin = os.Stdin
	agent.Stdout = os.Stdout
	agent.Debug = debug
	agent.plugins = make(map[string]Plugin)

	return agent
}
