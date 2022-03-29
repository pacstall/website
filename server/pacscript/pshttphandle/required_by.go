package pshttphandle

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/listener"
	"pacstall.dev/webserver/pacscript"
	"pacstall.dev/webserver/types"
	"pacstall.dev/webserver/types/list"
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

	pacpkg, err := allPackages.FindByName(name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	requiredBy := make([]*types.Pacscript, 0)
	for _, pkg := range allPackages.ToSlice() {
		if list.List[string](pacpkg.RequiredBy).Contains(list.Is(pkg.Name)) {
			requiredBy = append(requiredBy, pkg)
		}
	}

	listener.Json(w, requiredBy)
}
