package daemon

import (
	"github.com/kilgaloon/dongl/api"
	"github.com/spf13/viper"
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
	Config() *viper.Viper
	StartStop
	IsDebug() bool
	New(name string, cfg *viper.Viper, debug bool) Service
}
