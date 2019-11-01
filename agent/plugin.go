package agent

// Plugin is interface that represents basic plugin
type Plugin interface {
	Name() string
	Bootstrap() Plugin
}
