# Dongl

[![Build Status](https://travis-ci.com/kilgaloon/dongl.svg?branch=master)](https://travis-ci.com/kilgaloon/dongl)
[![codecov](https://codecov.io/gh/kilgaloon/dongl/branch/master/graph/badge.svg)](https://codecov.io/gh/kilgaloon/dongl)

Dongle strive to be boilerplate around your code and provide you with more control over your executing code, status and error reporting.

In order to start using Dongl you need to fullfill some of basic steps to wrap Dongl around you code.

First of all is creating `Agent`:

```

// Client settings and configurations
type Client struct {
	Name string
	*agent.Default
}

```
`Client` can be basically anything, and also can be `Potato`, this is your decission. We create struct to asign name which later will be used and also embed `agent.Default` which is default struct for agent, it provides you with basic commands and helpers that will help us register this agent as service later.

Now we provide our struct with other methods which are essential here `New, RName, Start and Stop`.

This is all about you, you define how your agent is started, stoped and what is his name for later distinction.

```

// New create client as a service
func (client *Client) New(name string, cfg *viper.Viper, debug bool) daemon.Service {
	a := agent.New(name, cfg, debug)
	c := &Client{
		name,
		a,
	}

	return c
}

// RName returns registered name for this service
func (client Client) RName() string {
	return client.Name
}

// Start client
func (client *Client) Start() {
	client.SetStatus("Started")
	fmt.Println("I Started!")

	go func() {
		for {
			if client.Status() == "Stopped" {
				fmt.Println("I Stopped!")
				break
			} else {
				fmt.Print("tick...")
			}

			time.Sleep(time.Second * 1)
		}
	}()
}

// Stop client
func (client *Client) Stop() {
	client.SetStatus("Stopped")
}
```

To enable commands we need to register api handles (keep in mind these are http endpoints).

These 2 commands provide us with 2 commands which can be used in following way: `./dongl --cmd="{agent} start"` and `./dongl --cmd="{agent} stop"`

Also we can pass arguments to our commands `./dongl --cmd="{agent} {cmd} {arg}..."`, these args are later available in http request in query as slice in order as your provided them in command.

So we can run echo command like `./dongl --cmd="{agent} echo hello"`

```

// RegisterAPIHandles to be used in http communication
func (client *Client) RegisterAPIHandles() map[string]api.Handle {
	cmds := make(map[string]api.Handle)

	cmds["stop"] = func(w http.ResponseWriter, r *http.Request) {
		client.Stop()
	}

	cmds["start"] = func(w http.ResponseWriter, r *http.Request) {
		client.Start()
	}

    cmds["echo"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Print(r.URL.Query()["args"][0])
	}

	return cmds
}
```