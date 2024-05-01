package repology_api

import (
	"fmt"
	"net/http"

	"pacstall.dev/webserver/services/server"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
)

func (c *RepologyController) GetRepologyPackageListHandle(w http.ResponseWriter, req *http.Request) error {
	etag := fmt.Sprintf("%v", c.packageCacheService.LastModified().UTC().String())
	if server.ApplyHeaders(c.serverConfiguration, etag, w, req) {
		// Response was cached and already sent
		return nil
	}

	results := c.findPackagesForRepology()

	server.Json(w, results)
	return nil
}

func (c *RepologyController) findPackagesForRepology() []repologyPackage {
	packages := c.packageCacheService.GetAll()

	packages = array.Filter(packages, func(it *array.Iterator[*pac.Script]) bool {
		return len(it.Value.Version) > 0
	})

	results := array.SwitchMap(packages, func(it *array.Iterator[*pac.Script]) repologyPackage {
		return newRepologyPackage(it.Value)
	})

	return results
}
