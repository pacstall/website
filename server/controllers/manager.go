package controllers

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/controller"
)

type ControllersManager struct {
	controllers []controller.Controller
	router      *mux.Router
}

func New(router *mux.Router, controllers []controller.Controller) *ControllersManager {
	m := &ControllersManager{}

	m.router = router
	m.controllers = controllers
	return m
}

func (m *ControllersManager) RegisterRoutes() {
	for _, controllerRef := range m.controllers {
		controllerName := reflect.TypeOf(controllerRef).Elem().Name()

		for _, route := range controllerRef.GetRoutes() {
			// This function wrap is necessary because otherwise the HandleFunc callback only takes the last item of the `routes` array
			// Passing the `route` as a parameter fixes this.

			m.router.HandleFunc(route.Path, func(route controller.ControllerRoute) func(w http.ResponseWriter, r *http.Request) {
				return func(w http.ResponseWriter, r *http.Request) {
					requestStart := time.Now()
					log.Debug("controller %v intercepted http request %v %v matching route %v", controllerName, route.Method, r.URL.String(), route.Path)

					if err := route.Handle(w, r); err != nil {
						log.Warn("http request %v %v failed with an internal server error. err: %+v", r.Method, r.URL.String(), err)
						w.WriteHeader(500)
					}

					duration := time.Since(requestStart)
					log.Debug("http request %v %v finished. total duration %v", route.Method, r.URL.String(), duration)
				}
			}(route)).Methods(string(route.Method))

			log.Info("controller %v registered route %v %v", controllerName, string(route.Method), route.Path)
		}
	}
}
