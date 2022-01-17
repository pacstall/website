package pacpkgs

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"pacstall.dev/website/serverlib"
)

func GetPackageListHandle(w http.ResponseWriter, req *http.Request) {

	packages := PackageList()
	page, err := strconv.ParseInt(req.URL.Query().Get("page"), 10, 32)
	if err != nil {
		page = 0
	}

	pageSize, err := strconv.ParseInt(req.URL.Query().Get("size"), 10, 32)
	if err != nil {
		pageSize = 50
	}

	etag := fmt.Sprintf("%v-%v-%v", LastModified().UTC().String(), page, pageSize)
	serverlib.ApplyCacheHeaders(etag, &w, req)

	startIndex := int(page * pageSize)
	endIndex := startIndex + int(pageSize)

	if len(packages) < startIndex {
		serverlib.SendJson(&w, []byte("[]"))
		return
	}

	if len(packages) < endIndex {
		endIndex = len(packages)
	}

	json, err := json.Marshal(packages[startIndex:endIndex])
	if err != nil {
		log.Printf("Could not marshal to json. Setting response 500.\n%v\n", err)
		w.WriteHeader(500)
	}

	serverlib.SendJson(&w, json)
}
