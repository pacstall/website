package server

import (
	"context"
	"errors"
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
var serverInstance *http.Server

func Router() *mux.Router {
	return &router
}

func Listen(port int) {
	registerHealthCheck()
	registerSiteMap()

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

		Router().PathPrefix("/").Handler(spaHandler{staticPath: path})
	}

	serverInstance = &http.Server{
		Handler:      Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := serverInstance.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		log.Info("Http server stopped")
	} else {
		log.Fatal("Could not start TCP listener on port %v. Got error: %v\n", port, err)
	}
}

func Shutdown() {
	if serverInstance == nil {
		log.Info("Server instance is already down")
	}

	ctx := context.Background()
	serverInstance.Shutdown(ctx)
	log.Info("Gracefully shutting down the http server")
}
