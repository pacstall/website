package psapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/server"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
	"pacstall.dev/webserver/types/pac/pacstore"
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

	found, err := allPackages.FindByName(name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	requiredBy := allPackages.Filter(func(p *pac.Script) bool {
		return list.List[string](found.RequiredBy).Contains(list.Is(p.Name))
	})

	server.Json(w, requiredBy)
}
