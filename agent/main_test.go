package agent

import (
	"testing"

	"github.com/kilgaloon/dongl/config"
)

var (
	iniFile      = "../tests/configs/config_regular.ini"
	path         = &iniFile
	cfgWrap      = config.NewConfigs()
	defaultAgent = New("test", cfgWrap.New("test", *path), false)
)

func TestGetterers(t *testing.T) {
	defaultAgent.RName()
	defaultAgent.Config()

	go defaultAgent.Write([]byte("test"))

	if defaultAgent.IsDebug() != defaultAgent.Debug {
		t.Fail()
	}

	defaultAgent.SetStatus("Stopped")
	if defaultAgent.Status() != "Stopped" {
		t.Fail()
	}

	h := defaultAgent.DefaultAPIHandles()
	if len(h) > 2 {
		t.Fail()
	}

}
