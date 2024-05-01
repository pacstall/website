package service

import (
	"github.com/gorilla/mux"
)

type ServerService interface {
	Listen()
	Router() *mux.Router
	Shutdown()
}
