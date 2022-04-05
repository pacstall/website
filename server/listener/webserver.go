package listener

import (
	"fmt"
	"net/http"
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

func Listen(port uint16) {
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

	if config.IsProduction {
		Router().PathPrefix("/").Handler(spaHandler{staticPath: config.TCPServer.PublicDir})
	}

	server := &http.Server{
		Handler:      Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	log.Error.Fatalf("Could not start TCP listener on port %v. Got error: %v\n", port, err)
}
