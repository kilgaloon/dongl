package api

import (
	"net/http"
)

// Handle is type that defines handler of http request
type Handle func(w http.ResponseWriter, r *http.Request)

var mux = http.NewServeMux()

// API defines socket on which we listen for commands
type API struct {
	HTTP *http.Server
}

// Registrator defines interface that help us to register http handler
type Registrator interface {
	RName() string
	RegisterAPIHandles() map[string]Handle
	DefaultAPIHandles() map[string]Handle
}

// Register all available http handles
func (a *API) Register(r Registrator) *API {
	for p, f := range r.RegisterAPIHandles() {
		mux.HandleFunc("/"+r.RName()+"/"+p, f)
	}

	for p, f := range r.DefaultAPIHandles() {
		mux.HandleFunc("/"+r.RName()+"/"+p, f)
	}

	return a
}

// RegisterHandle append handle before server is started
func (a *API) RegisterHandle(e string, h Handle) {
	mux.HandleFunc("/"+e, h)
}

// Start api server
func (a *API) Start() {
	a.HTTP.Handler = mux
	if err := a.HTTP.ListenAndServe(); err != nil {
		panic(err)
	}
}

// New creates new socket
func New() *API {
	api := &API{
		&http.Server{
			Addr: ":11401",
		},
	}

	return api
}
