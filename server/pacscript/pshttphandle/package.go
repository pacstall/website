package pshttphandle

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/listener"
	"pacstall.dev/webserver/pacscript"
)

func GetPacscriptHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]
	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	if listener.ApplyHeaders(fmt.Sprintf("%v-%v", pacscript.LastModified().UTC().String(), name), w, req) {
		return // req is cached
	}

	pkg, err := pacscript.GetAll().FindByName(name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	listener.Json(w, pkg)
}
