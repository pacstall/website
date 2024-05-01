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
	"pacstall.dev/webserver/types/service"
)

type WebserverService struct {
	router              *mux.Router
	server              *http.Server
	serverConfiguration config.ServerConfiguration
	packageCacheService service.PackageCacheService
}

// Initializes the given service. If nil, it creates a new one.
func New(serverConfiguration config.ServerConfiguration, packageCacheService service.PackageCacheService) *WebserverService {
	service := &WebserverService{}

	service.router = mux.NewRouter()
	service.serverConfiguration = serverConfiguration
	service.packageCacheService = packageCacheService
	service.server = &http.Server{
		Handler:      service.router,
		Addr:         fmt.Sprintf("0.0.0.0:%v", serverConfiguration.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return service
}

func (s *WebserverService) Router() *mux.Router {
	return s.router
}

func (s *WebserverService) Listen() {
	s.registerHealthCheck()
	s.registerSiteMap()

	s.router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("X-Pacstall-Version", s.serverConfiguration.Version)

			if strings.Contains(r.URL.Path, "/api") {
				w.Header().Add("Content-Type", "application/json")
			}
			next.ServeHTTP(w, r)
		})
	})

	go triggerServerOnline(s.serverConfiguration.Port)

	if s.serverConfiguration.Production {
		path, err := filepath.Abs(s.serverConfiguration.PublicDir)
		if err != nil {
			log.Fatal("failed to find client public dir at path '%s'. err: %+v", s.serverConfiguration.PublicDir, err)
		}

		s.router.PathPrefix("/").Handler(spaHandler{staticPath: path})
	}

	err := s.server.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		log.Info("http server stopped")
	} else {
		log.Fatal("could not start TCP listener on port %v. Got error: %+v\n", s.serverConfiguration.Port, err)
	}
}

func (s *WebserverService) Shutdown() {
	if s.server == nil {
		log.Info("server instance is already down")
	}

	ctx := context.Background()
	s.server.Shutdown(ctx)
	log.Info("gracefully shutting down the http server")
}
