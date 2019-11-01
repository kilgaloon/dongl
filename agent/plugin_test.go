package agent

import (
	"testing"
)

// Worker do what i say
type TestPlugin struct {
	booted bool
}

//Name of the plugin
func (p TestPlugin) Name() string {
	return "TestPlugin"
}

//Bootstrap the plugin
func (p TestPlugin) Bootstrap() Plugin {
	plugin := &TestPlugin{}

	plugin.booted = true

	return plugin
}

func TestPluginRegistrationAndBootstrap(t *testing.T) {
	defaultAgent.RegisterPlugin(TestPlugin{})

	p := defaultAgent.Plugin("TestPlugin")
	if p.Name() != "TestPlugin" {
		t.Fail()
	}
}
