package main

import (
	"github.com/kilgaloon/dongl/daemon"
	"github.com/kilgaloon/dongl/services/example"
)

func main() {
	// add your services
	daemon.Srv.AddService(&example.Client{Name: "example_client"})

	daemon.Srv.Run()
}
