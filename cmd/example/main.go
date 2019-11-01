package main

import (
	"github.com/kilgaloon/dongl/daemon"
	"github.com/kilgaloon/dongl/services/example"
)

func main() {
	// init daemon
	srv := daemon.Init()

	// Add your services to daemon
	srv.AddService(&example.Client{Name: "example_client"})

	// run daemon
	srv.Run()
}
