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

func GetPacscriptHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]
	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	if server.ApplyHeaders(fmt.Sprintf("%v-%v", pacstore.LastModified().UTC().String(), name), w, req) {
		return // req is cached
	}

	pkg, err := array.FindBy(pacstore.GetAll(), func(s *pac.Script) bool {
		return s.PackageName == name
	})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	server.Json(w, pkg)
}
