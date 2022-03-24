package pshttphandle

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"pacstall.dev/website/listener"
	"pacstall.dev/website/pacscript"
	"pacstall.dev/website/types"
)

func GetPackagesRequiredByHandle(w http.ResponseWriter, req *http.Request) {
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

	allPackages := pacscript.PackageList()

	var pacpkg *types.PackageInfo
	for _, it := range allPackages {
		if name == it.Name {
			pacpkg = it
			break
		}
	}

	if pacpkg == nil {
		w.WriteHeader(404)
		return
	}

	requiredBy := make([]*types.PackageInfo, 0)
	for _, pkg := range allPackages {
		for _, requiredByName := range pacpkg.RequiredBy {
			if strings.Compare(pkg.Name, requiredByName) == 0 {
				requiredBy = append(requiredBy, pkg)
				break
			}
		}
	}

	listener.Json(w, requiredBy)
}
