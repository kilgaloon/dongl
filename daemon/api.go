package daemon

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/kilgaloon/dongl/api"
	"github.com/olekukonko/tablewriter"
)

// InfoResponse defines how response looks like when
// info for running daemon is returned
type InfoResponse struct {
	PID        int
	ConfigPath string
	PidPath    string
	Debug      bool
	Memory     string
}

// ServicesListResponse  defines how response looks like when
// list of services is requested
type ServicesListResponse struct {
	List [][]string
}

// GetInfo display info for running daemon
func (d *Daemon) GetInfo() *InfoResponse {
	r, err := api.HTTPClient.Get(api.RevealEndpoint("/info", api.Cmd("info")))
	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	resp := &InfoResponse{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		log.Fatal(err)
	}

	return resp
}

func (d *Daemon) renderInfo() {
	table := tablewriter.NewWriter(os.Stdout)

	resp := d.GetInfo()

	pid := strconv.Itoa(resp.PID)
	debug := "No"
	if resp.Debug {
		debug = "Yes"
	}

	table.SetHeader([]string{"PID", "Config path", "Pid path", "Debug", "Memory"})
	table.Append([]string{pid, resp.ConfigPath, resp.PidPath, debug, resp.Memory})

	table.Render()
}
