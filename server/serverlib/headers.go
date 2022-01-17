package serverlib

import (
	"net/http"
	"strings"
)

type AlreadyResponded = bool

func ApplyCacheHeaders(etag string, w *http.ResponseWriter, r *http.Request) AlreadyResponded {
	(*w).Header().Add("Cache-Control", "max-age:420") // 7 minutes
	(*w).Header().Add("Etag", etag)

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			(*w).WriteHeader(http.StatusNotModified)
			return true
		}
	}

	return false
}

func SendJson(w *http.ResponseWriter, data []byte) {
	(*w).Header().Add("Content-Type", "application/json")
	(*w).Write(data)
}
