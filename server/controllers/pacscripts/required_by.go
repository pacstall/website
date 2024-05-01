package psapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/services/server"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

func (c *PackageController) GetPackageRequiredByHandle(w http.ResponseWriter, req *http.Request) error {
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

	allPackages := c.packageCacheService.GetAll()

	found, err := array.FindBy(allPackages, func(p *pac.Script) bool {
		return p.PackageName == name
	})

	if err != nil {
		w.WriteHeader(404)
		return nil
	}

	requiredBy := array.Filter(allPackages, func(it *array.Iterator[*pac.Script]) bool {
		return array.Contains(found.RequiredBy, array.Is(it.Value.PackageName))
	})

	server.Json(w, requiredBy)
	return nil
}
