package psapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/services/server"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

type pacscriptDependencies struct {
	RuntimeDependencies  []string      `json:"runtimeDependencies"`
	BuildDependencies    []string      `json:"buildDependencies"`
	OptionalDependencies []string      `json:"optionalDependencies"`
	PacstallDependencies []*pac.Script `json:"pacstallDependencies"`
}

func (c *PackageController) GetPackageDependenciesHandle(w http.ResponseWriter, req *http.Request) error {
	params := mux.Vars(req)
	name, ok := params["name"]

	etag := fmt.Sprintf("%v-%v", c.packageCacheService.LastModified().UTC().String(), name)
	if server.ApplyHeaders(c.serverConfiguration, etag, w, req) {
		// Response was cached and already sent
		return nil
	}

	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return nil
	}

	dependencies, ok := c.findPackageDependencies(name)
	if !ok {
		w.WriteHeader(404)
		return nil
	}

	server.Json(w, dependencies)
	return nil
}

func (c *PackageController) findPackageDependencies(name string) (pacscriptDependencies, bool) {
	allPackages := c.packageCacheService.GetAll()

	foundPackage, err := array.FindBy(allPackages, func(s *pac.Script) bool {
		return s.PackageName == name
	})

	if err != nil {
		return pacscriptDependencies{}, false
	}

	pacstallDependencies := make([]*pac.Script, 0)
	for _, pkg := range foundPackage.PacstallDependencies {
		if found, err := array.FindBy(allPackages, func(pi *pac.Script) bool { return pkg == pi.PackageName }); err == nil {
			pacstallDependencies = append(pacstallDependencies, found)
		} else {
			log.Error("could not find pacstall dependency %s of package %s.\n", pkg, foundPackage.PackageName)
		}
	}

	return pacscriptDependencies{
		RuntimeDependencies:  foundPackage.RuntimeDependencies,
		BuildDependencies:    foundPackage.BuildDependencies,
		OptionalDependencies: foundPackage.OptionalDependencies,
		PacstallDependencies: pacstallDependencies,
	}, true
}
