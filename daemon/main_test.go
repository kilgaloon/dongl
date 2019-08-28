package daemon

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/kilgaloon/dongl/api"
	"github.com/kilgaloon/dongl/config"
)

var (
	configs            = config.NewConfigs()
	ConfigWithSettings = configs.New("test", "../tests/configs/config_regular.ini")
)

type fakeService struct {
}

func (fs fakeService) RName() string {
	return "fake_service"
}

// RegisterAPIHandles to be used in socket communication
// If you want to takeover default commands from agent
// call DefaultCommands from Agent which is same command
func (fs *fakeService) RegisterAPIHandles() map[string]api.Handle {
	cmds := make(map[string]api.Handle)

	return cmds
}

func (fs *fakeService) DefaultAPIHandles() map[string]api.Handle {
	cmds := make(map[string]api.Handle)

	return cmds
}

func (fs *fakeService) Status() string {
	return "Testing"
}

func (fs *fakeService) Config() config.AgentConfig {
	return ConfigWithSettings
}

func (fs *fakeService) IsDebug() bool {
	return true
}

func (fs *fakeService) SetStatus(s string)         {}
func (fs *fakeService) Start()                  {}
func (fs *fakeService) Stop()                   {}
func (fs *fakeService) Pause()                  {}
func (fs *fakeService) SetPipeline(chan string) {}
func (fs *fakeService) New(name string, cfg config.AgentConfig, debug bool) Service {
	srv := &fakeService{}
	return srv
}

func TestAddService(t *testing.T) {
	fk := &fakeService{}
	Srv.AddService(fk)
}

func TestRun(t *testing.T) {
	go Srv.Run()

	if Srv.PID() != os.Getpid() {
		t.Fatal("PID NOT MATCHED")
	}

	req, err := http.NewRequest("GET", "/info", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	Srv.daemonInfo(rr, req)

	for i := 0; i < 5; i++ {
		_, err := http.Get("http://localhost:11401")
		if err != nil {
			// handle error
			time.Sleep(2 * time.Second)
			continue
		}

		Srv.GetInfo()
		Srv.renderInfo()

	}

	//Srv.Kill()
}

func TestRunningDaemonInfo(t *testing.T) {
	Srv.Cmd = "info"
	Srv.Run()
	for i := 0; i < 5; i++ {
		_, err := http.Get("http://localhost:11401")
		if err != nil {
			// handle error
			Srv.API.Start()
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}
}

func TestRunningDaemonServices(t *testing.T) {
	Srv.Cmd = "services"
	Srv.Run()
	for i := 0; i < 5; i++ {
		_, err := http.Get("http://localhost:11401")
		if err != nil {
			// handle error
			Srv.API.Start()
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}
}

func TestRunningDaemonKill(t *testing.T) {
	Srv.Cmd = "kill"
	Srv.Run()
	for i := 0; i < 5; i++ {
		_, err := http.Get("http://localhost:11401")
		if err != nil {
			// handle error
			Srv.API.Start()
			time.Sleep(2 * time.Second)
			continue
		}

		break
	}
}
