package listener

import (
	"encoding/json"
	"net/http"
	"strings"

	"pacstall.dev/webserver/config"
)

type AlreadyResponded = bool

func ApplyHeaders(etag string, w http.ResponseWriter, r *http.Request) AlreadyResponded {
	w.Header().Add("Cache-Control", "max-age:420") // 7 minutes
	w.Header().Add("Etag", etag)

	if match := r.Header.Get("If-None-Match"); match != "" {
		if strings.Contains(match, etag) {
			w.WriteHeader(http.StatusNotModified)
			return true
		}
	}

	if !config.Config.Production {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:1234")
		w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	}

	return false
}

func SendJson(w http.ResponseWriter, data []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func Json(w http.ResponseWriter, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	SendJson(w, data)
}
