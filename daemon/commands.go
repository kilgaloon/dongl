package daemon

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"strconv"

	"github.com/kilgaloon/dongl/api"
)

func (d *Daemon) daemonInfo(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	alloc := strconv.FormatFloat(float64(mem.Alloc/1024)/1024, 'f', 2, 64)

	resp := &InfoResponse{
		PID:        d.PID(),
		ConfigPath: d.ConfigPath(),
		PidPath:    d.PidPath(),
		Debug:      d.Debug,
		Memory:     alloc + "MiB",
	}

	w.WriteHeader(http.StatusOK)

	j, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(j)

	return
}

func (d *Daemon) daemonKill(w http.ResponseWriter, r *http.Request) {
	resp := api.TableResponse{
		Header:  []string{"Message"},
		Columns: [][]string{},
	}

	d.Kill()
	if api.IsAPIRunning() {
		resp.Columns = append(resp.Columns, []string{"Failed to kill daemon"})
	} else {
		resp.Columns = append(resp.Columns, []string{"Daemon killed"})
	}

	j, err := json.Marshal(resp)
	if err != nil {
		resp.Columns = append(resp.Columns, []string{"Daemon killed"})
	}

	w.Write(j)
}

func (d *Daemon) servicesList(w http.ResponseWriter, r *http.Request) {
	resp := api.TableResponse{
		Header:  []string{"Agent name", "Status"},
		Columns: [][]string{},
	}

	for agent, service := range d.services {
		resp.Columns = append(resp.Columns, []string{agent, service.Status()})
	}

	w.WriteHeader(http.StatusOK)

	j, err := json.Marshal(resp)
	if err != nil {
		log.Fatal(err)
	}

	w.Write(j)

	return
}
