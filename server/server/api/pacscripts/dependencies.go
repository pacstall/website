package psapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/server"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
)

type pacscriptDependencies struct {
	RuntimeDependencies  []string      `json:"runtimeDependencies"`
	BuildDependencies    []string      `json:"buildDependencies"`
	OptionalDependencies []string      `json:"optionalDependencies"`
	PacstallDependencies []*pac.Script `json:"pacstallDependencies"`
}

func GetPacscriptDependenciesHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]

	etag := fmt.Sprintf("%v-%v", pacstore.LastModified().UTC().String(), name)
	if server.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	allPacscripts := pacstore.GetAll()

	pacpkg, err := allPacscripts.FindByName(name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	pacstallDependencies := make([]*pac.Script, 0)
	for _, pkg := range pacpkg.PacstallDependencies {
		if found, err := pacstore.GetAll().FindBy(func(pi *pac.Script) bool { return pkg == pi.Name }); err == nil {
			pacstallDependencies = append(pacstallDependencies, found)
		} else {
			log.Error("Could not find pacstall dependency %s of package %s.\n", pkg, pacpkg.Name)
		}
	}

	response := pacscriptDependencies{
		RuntimeDependencies:  pacpkg.RuntimeDependencies,
		BuildDependencies:    pacpkg.BuildDependencies,
		OptionalDependencies: pacpkg.OptionalDependencies,
		PacstallDependencies: pacstallDependencies,
	}

	server.Json(w, response)
}
