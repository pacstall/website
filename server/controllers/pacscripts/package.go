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

func (c *PackageController) GetPackageHandle(w http.ResponseWriter, req *http.Request) error {
	log.Info("1\n")

	params := mux.Vars(req)
	name, ok := params["name"]
	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return nil
	}

	log.Info("2\n")
	if server.ApplyHeaders(c.serverConfiguration, fmt.Sprintf("%v-%v", c.packageCacheService.LastModified().UTC().String(), name), w, req) {

		return nil // req is cached
	}

	log.Info("3\n")
	pkg, err := array.FindBy(c.packageCacheService.GetAll(), func(s *pac.Script) bool {
		return s.PackageName == name
	})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}

	log.Info("4\n")
	log.Info("found: %v\n", pkg)

	server.Json(w, pkg)
	return nil
}
