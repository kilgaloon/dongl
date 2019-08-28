package daemon

import (
	"github.com/kilgaloon/dongl/api"
	"github.com/kilgaloon/dongl/config"
)

// StartStop defines service that can be started and stoped
type StartStop interface {
	Start()
	Stop()
}

// Service struct define
type Service interface {
	api.Registrator
	Status() string
	SetStatus(s string)
	Config() config.AgentConfig
	StartStop
	IsDebug() bool
	New(name string, cfg config.AgentConfig, debug bool) Service
}
