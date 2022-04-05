package repology_api

import (
	"fmt"
	"net/http"

	"pacstall.dev/webserver/listener"
	"pacstall.dev/webserver/store/pacstore"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

func GetRepologyPackageListHandle(w http.ResponseWriter, req *http.Request) {
	packages := pacstore.GetAll().ToSlice()

	etag := fmt.Sprintf("%v", pacstore.LastModified().UTC().String())
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	results := list.Map(packages, func(_ int, p *pac.Script) repologyPackage {
		return newRepologyPackage(*p)
	})

	listener.Json(w, results)
}
