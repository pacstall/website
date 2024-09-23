package psapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/internal/pacnexus/server"
	"pacstall.dev/webserver/internal/pacnexus/types/pac"
	"pacstall.dev/webserver/internal/pacnexus/types/pac/pacstore"
	"pacstall.dev/webserver/pkg/common/array"
)

type pacscriptDependencies struct {
	RuntimeDependencies  []pac.ArchDistroString `json:"runtimeDependencies"`
	BuildDependencies    []pac.ArchDistroString `json:"buildDependencies"`
	OptionalDependencies []pac.ArchDistroString `json:"optionalDependencies"`
	PacstallDependencies []pac.ArchDistroString `json:"pacstallDependencies"`
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

	pacpkg, err := array.FindBy(allPacscripts, func(s *pac.Script) bool {
		return s.PackageName == name
	})

	if err != nil {
		w.WriteHeader(404)
		return
	}

	response := pacscriptDependencies{
		RuntimeDependencies:  pacpkg.RuntimeDependencies,
		BuildDependencies:    pacpkg.BuildDependencies,
		OptionalDependencies: pacpkg.OptionalDependencies,
		PacstallDependencies: pacpkg.PacstallDependencies,
	}

	server.Json(w, response)
}
