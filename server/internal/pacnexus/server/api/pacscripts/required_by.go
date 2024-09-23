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

func GetPacscriptRequiredByHandle(w http.ResponseWriter, req *http.Request) {
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

	allPackages := pacstore.GetAll()

	found, err := array.FindBy(allPackages, func(p *pac.Script) bool {
		return p.PackageName == name
	})

	if err != nil {
		w.WriteHeader(404)
		return
	}

	requiredBy := array.Filter(allPackages, func(it *array.Iterator[*pac.Script]) bool {
		return array.Contains(found.RequiredBy, array.Is(it.Value.PackageName))
	})

	server.Json(w, requiredBy)
}
