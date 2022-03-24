package pshttphandle

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"pacstall.dev/website/listener"
	"pacstall.dev/website/pacscript"
)

func GetPackageHandle(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	name, ok := params["name"]
	if !ok || len(name) < 2 {
		w.WriteHeader(400)
		return
	}

	if listener.ApplyHeaders(fmt.Sprintf("%v-%v", pacscript.LastModified().UTC().String(), name), w, req) {
		return // req is cached
	}

	for _, pkg := range pacscript.PackageList() {
		if strings.Compare(pkg.Name, name) == 0 {
			listener.Json(w, pkg)
			return
		}
	}

	w.WriteHeader(404)
}
