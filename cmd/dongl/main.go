package main

import (
	"github.com/kilgaloon/dongl/daemon"
	"github.com/kilgaloon/dongl/services/example"
)

func main() {
	// add your services
	srv := daemon.Init()
	
	srv.AddService(&example.Client{Name: "example_client"})

	srv.Run()
}
