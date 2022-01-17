package serverlib

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"pacstall.dev/website/cfg"
)

var router mux.Router = *mux.NewRouter()

func Router() *mux.Router {
	return &router
}

func Serve(port int) {
	registerHealthCheck()

	Router().Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, "/api") {
				w.Header().Add("Content-Type", "application/json")
			}
			next.ServeHTTP(w, r)
		})
	})

	go triggerServerOnline(port)

	if cfg.Config.Production {
		Router().PathPrefix("/").Handler(spaHandler{staticPath: cfg.Config.TCPServer.PublicDir, indexPath: "index.html"})
	}

	server := &http.Server{
		Handler:      Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	log.Panicf("Could not start TCP listener on port %v. Got error: %v", port, err)
}
