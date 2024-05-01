package psapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/services/server"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

func (c *PackageController) GetPackageHandle(w http.ResponseWriter, req *http.Request) error {
	params := mux.Vars(req)
	name, ok := params["name"]
	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return nil
	}

	if server.ApplyHeaders(c.serverConfiguration, fmt.Sprintf("%v-%v", c.packageCacheService.LastModified().UTC().String(), name), w, req) {
		return nil // req is cached
	}

	pkg, ok := c.findPackageByName(name)

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	server.Json(w, pkg)
	return nil
}

func (c *PackageController) findPackageByName(name string) (*pac.Script, bool) {
	pkg, err := array.FindBy(c.packageCacheService.GetAll(), func(s *pac.Script) bool {
		return s.PackageName == name
	})

	return pkg, err == nil
}
