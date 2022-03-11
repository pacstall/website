package pshttphandle

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/website/listener"
	"pacstall.dev/website/pacscript"
	"pacstall.dev/website/types"
)

type packageDependencies struct {
	RuntimeDependencies  []string             `json:"runtimeDependencies"`
	BuildDependencies    []string             `json:"buildDependencies"`
	OptionalDependencies []string             `json:"optionalDependencies"`
	PacstallDependencies []*types.PackageInfo `json:"pacstallDependencies"`
}

func GetPackageDependenciesHandle(w http.ResponseWriter, req *http.Request) {
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

	allPackages := pacscript.PackageList()

	var pacpkg *types.PackageInfo
	for _, it := range allPackages {
		if name == it.Name {
			pacpkg = it
			break
		}
	}

	if pacpkg == nil {
		w.WriteHeader(404)
		return
	}

	pacstallDependencies := make([]*types.PackageInfo, 0)
	for _, pkg := range pacpkg.PacstallDependencies {
		if found := pacscript.FindPackageInList(pkg, allPackages); found != nil {
			pacstallDependencies = append(pacstallDependencies, found)
		} else {
			log.Printf("Could not find pacstall dependency %s of package %s.\n", pkg, pacpkg.Name)
		}
	}

	response := packageDependencies{
		RuntimeDependencies:  pacpkg.RuntimeDependencies,
		BuildDependencies:    pacpkg.BuildDependencies,
		OptionalDependencies: pacpkg.OptionalDependencies,
		PacstallDependencies: pacstallDependencies,
	}

	listener.Json(w, response)
}
