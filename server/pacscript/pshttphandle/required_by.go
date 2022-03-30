package pshttphandle

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/listener"
	"pacstall.dev/webserver/pacscript"
	"pacstall.dev/webserver/types/list"
	"pacstall.dev/webserver/types/pac"
)

func GetPacscriptRequiredByHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]

	etag := fmt.Sprintf("%v-%v", pacscript.LastModified().UTC().String(), name)
	if listener.ApplyHeaders(etag, w, req) {
		// Response was cached and already sent
		return
	}

	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	allPackages := pacscript.GetAll()

	found, err := allPackages.FindByName(name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	requiredBy := allPackages.Filter(func(p *pac.Script) bool {
		return list.List[string](found.RequiredBy).Contains(list.Is(p.Name))
	})

	listener.Json(w, requiredBy)
}
