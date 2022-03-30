package pshttphandle

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/listener"
	"pacstall.dev/webserver/pacscript"
	"pacstall.dev/webserver/types/pac"
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

	etag := fmt.Sprintf("%v-%v", pacscript.LastModified().UTC().String(), name)
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	allPacscripts := pacscript.GetAll()

	pacpkg, err := allPacscripts.FindByName(name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	pacstallDependencies := make([]*pac.Script, 0)
	for _, pkg := range pacpkg.PacstallDependencies {
		if found, err := pacscript.GetAll().FindBy(func(pi *pac.Script) bool { return pkg == pi.Name }); err == nil {
			pacstallDependencies = append(pacstallDependencies, found)
		} else {
			log.Printf("Could not find pacstall dependency %s of package %s.\n", pkg, pacpkg.Name)
		}
	}

	response := pacscriptDependencies{
		RuntimeDependencies:  pacpkg.RuntimeDependencies,
		BuildDependencies:    pacpkg.BuildDependencies,
		OptionalDependencies: pacpkg.OptionalDependencies,
		PacstallDependencies: pacstallDependencies,
	}

	listener.Json(w, response)
}
