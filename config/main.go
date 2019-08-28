package config

import (
	"path/filepath"

	"gopkg.in/ini.v1"
)

// Configs for different agents
type Configs struct {
	cfgs map[string]AgentConfig
}

// NewConfigs return Configs of agents
func NewConfigs() *Configs {
	return &Configs{
		cfgs: make(map[string]AgentConfig),
	}
}

//Config return config by name of the agent
func (c *Configs) Config(name string) AgentConfig {
	if cfg, ok := c.cfgs[name]; ok {
		return cfg
	}

	return AgentConfig{}
}

// AgentConfig holds config for agents
type AgentConfig struct {
	path   string
	values map[string]string
}

// Path returns path of config file
func (ac AgentConfig) Path() string {
	p, err := filepath.Abs(ac.path)
	if err != nil {
		return ac.path
	}

	return p
}

// Value returns path of config file
func (ac AgentConfig) Value(v string) string {
	if value, ok := ac.values[v]; ok {
		return value
	}

	return ""
}

// New Create new config
func (c *Configs) New(name string, path string) AgentConfig {
	// more config processors here
	cfg, err := ini.Load(path)
	if err != nil {
		panic(err)
	}

	ac := AgentConfig{}
	ac.values = make(map[string]string)
	ac.path = path

	for _, key := range cfg.Section("").Keys() {
		ac.values[key.Name()] = key.Value()
	}

	c.cfgs[name] = ac
	return ac
}
