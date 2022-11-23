package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
)

var router mux.Router = *mux.NewRouter()

func Router() *mux.Router {
	return &router
}

func Listen(port int) {
	registerHealthCheck()

	Router().Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Pacstall-Version", config.Version)

			if strings.Contains(r.URL.Path, "/api") {
				w.Header().Add("Content-Type", "application/json")
			}
			next.ServeHTTP(w, r)
		})
	})

	go triggerServerOnline(port)

	if config.Production {
		path, err := filepath.Abs(config.PublicDir)
		if err != nil {
			log.Fatal("failed to find client public dir at path '%s'. err: %v", config.PublicDir, err)
		}

		Router().PathPrefix("/").Handler(spaHandler{ path })
	}

	server := &http.Server{
		Handler:      Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	log.Fatal("Could not start TCP listener on port %v. Got error: %v\n", port, err)
}
