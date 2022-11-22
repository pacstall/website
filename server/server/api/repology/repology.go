package repology_api

import (
	"fmt"
	"net/http"

	"pacstall.dev/webserver/server"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
)

func GetRepologyPackageListHandle(w http.ResponseWriter, req *http.Request) {
	packages := pacstore.GetAll().Filter(func(s *pac.Script) bool {
		return len(s.Version) > 0
	}).ToSlice()

	etag := fmt.Sprintf("%v", pacstore.LastModified().UTC().String())
	if server.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	results := list.Map(packages, func(_ int, p *pac.Script) repologyPackage {
		return newRepologyPackage(*p)
	})

	server.Json(w, results)
}
