package example

import "github.com/kilgaloon/dongl/agent"

// Worker do what i say
type Worker struct {
}

//Name of the plugin
func (w Worker) Name() string {
	return "Worker"
}

//Bootstrap the plugin
func (w Worker) Bootstrap() agent.Plugin {
	return &Worker{}
}
