package serverlib

import (
	"fmt"
	"log"
	"net/http"
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

	if cfg.Config.Production {
		Router().PathPrefix("/").Handler(http.FileServer(http.Dir(cfg.Config.TCPServer.PublicDir)))

	}

	go triggerServerOnline(port)

	server := &http.Server{
		Handler:      Router(),
		Addr:         fmt.Sprintf("0.0.0.0:%v", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()
	log.Panicf("Could not start TCP listener on port %v. Got error: %v", port, err)
}
