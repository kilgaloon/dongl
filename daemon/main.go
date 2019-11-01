package daemon

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kilgaloon/dongl/api"
	"github.com/spf13/viper"
)

// Daemon is long living process that serves as middleware
// and access to multiple agents
type Daemon struct {
	pid          int
	pidPath      string
	PidFile      *os.File
	services     map[string]Service
	Config       *viper.Viper
	configPath   string
	Cmd          api.Cmd
	Debug        bool
	API          *api.API
	shutdownChan chan bool
}

// PID gets current PID of client
func (d *Daemon) PID() int {
	return d.pid
}

// ConfigPath returns path of config file
func (d *Daemon) ConfigPath() string {
	p, err := filepath.Abs(d.configPath)
	if err != nil {
		return d.configPath
	}

	return p
}

// PidPath returns path of config file
func (d *Daemon) PidPath() string {
	p, err := filepath.Abs(d.pidPath)
	if err != nil {
		return d.pidPath
	}

	return p
}

// AddService push agent as a service to list of services
func (d *Daemon) AddService(s Service) {
	var cfg *viper.Viper

	name := s.RName()

	viper.SetConfigName(name)
	viper.AddConfigPath(d.configPath)
	err := viper.ReadInConfig()
	// if there is error reading specific file for service load global config
	// otherwise read specific config file for service
	if err != nil {            
		cfg = d.Config
	} else {
		cfg = viper.GetViper()
	}

	a := s.New(name, cfg, d.Debug)

	d.API.Register(a)

	d.services[name] = a
}

// Run starts daemon and long living process
func (d *Daemon) Run() {
	if api.IsAPIRunning() {
		// more commands can/will be used here

		switch d.Cmd.Agent() {
		case "info":
			d.renderInfo()
			return
		}

		api.Resolver(d.Cmd)
	} else {
		go func() {
			for _, s := range d.services {
				log.Printf("Starting service %s", s.RName())
				go s.Start()

				break
			}

			d.API.RegisterHandle("info", d.daemonInfo)
			d.API.RegisterHandle("kill", d.daemonKill)
			d.API.RegisterHandle("services", d.servicesList)
			d.API.Start()
		}()

		for {
			select {
			case <-d.shutdownChan:
				if os.Getenv("RUN_MODE") != "test" {
					os.Exit(1)
				}
				break
			}
		}

	}
}

// Kill daemon and remove .pid file
func (d *Daemon) Kill() {
	err := os.Remove(d.PidPath())
	if err != nil {
		panic(err)
	}

	d.shutdownChan <- true
}

// Init -ialize dameon
func Init() *Daemon {
	var configPath, pidPath *string
	var debug *bool
	var pid int

	d := new(Daemon)
	if api.IsAPIRunning() {
		resp := d.GetInfo()

		configPath = &resp.ConfigPath
		pidPath = &resp.PidPath
		debug = &resp.Debug
		pid = resp.PID
	} else {
		if os.Getenv("RUN_MODE") == "test" {
			pp := "../tests/var/run/dongl/.pid"
			cp := "../tests/configs"
			dbg := true

			pidPath = &pp
			configPath = &cp
			debug = &dbg
		} else {
			configPath = flag.String("cfg", "/etc/dongl", "Path to .ini configuration")
			pidPath = flag.String("pid", "/var/run/dongl/.pid", "PID file of process")
			debug = flag.Bool("debug", false, "Debug mode")
		}
	}

	cmd := flag.String("cmd", "info", "Shows basic information about daemon")
	// https://golang.org/doc/go1.13#testing
	if os.Getenv("RUN_MODE") != "test" {
		flag.Parse()
	}

	f, err := os.OpenFile(*pidPath, os.O_RDWR|os.O_CREATE, 0644)
	d.PidFile = f
	d.pidPath = *pidPath
	if err != nil {
		panic(err)
	}

	if pid == 0 {
		d.pid = os.Getpid()
		pid := strconv.Itoa(d.pid)
		_, err = d.PidFile.WriteString(pid)
		if err != nil {
			log.Fatal("Failed to start client, can't save PID")
		}
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(*configPath)
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}

	d.configPath = *configPath
	d.Config = viper.GetViper()

	d.services = make(map[string]Service)
	d.Debug = *debug
	d.Cmd = api.Cmd(*cmd)
	d.API = api.New()
	d.shutdownChan = make(chan bool, 1)

	return d
}
