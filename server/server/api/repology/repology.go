package repology_api

import (
	"fmt"
	"net/http"

	"pacstall.dev/webserver/server"
	"pacstall.dev/webserver/types/array"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
)

func GetRepologyPackageListHandle(w http.ResponseWriter, req *http.Request) {
	packages := pacstore.GetAll()

	packages = array.Filter(packages, func(it *array.Iterator[*pac.Script]) bool {
		return len(it.Value.Version) > 0
	})

	etag := fmt.Sprintf("%v", pacstore.LastModified().UTC().String())
	if server.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	results := array.SwitchMap(packages, func(it *array.Iterator[*pac.Script]) repologyPackage {
		return newRepologyPackage(it.Value)
	})

	server.Json(w, results)
}
