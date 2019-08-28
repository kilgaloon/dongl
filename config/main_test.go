package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configs                      = NewConfigs()
	ConfigWithSettings           = configs.New("test", "../tests/configs/config_regular.ini")
)

func TestBuildWithSettings(t *testing.T) {
	cfg := configs.Config("test")

	p, _ := filepath.Abs("../tests/configs/config_regular.ini")
	assert.Equal(t, p, cfg.Path())
	assert.Equal(t, cfg.Value("log"), "../../tests/var/log/error.log")
}