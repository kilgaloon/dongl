package agent

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

var (
	defaultAgent = BuildDefaultAgent()
)

func BuildDefaultAgent() *Default {
	viper.SetConfigName("config")
	viper.AddConfigPath("../tests/configs")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	return New("test", viper.GetViper(), false)
}

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
