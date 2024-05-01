package controller

import "net/http"

type method string

const (
	METHOD_GET     method = "GET"
	METHOD_POST    method = "POST"
	METHOD_PUT     method = "PUT"
	METHOD_DELETE  method = "DELETE"
	METHOD_OPTIONS method = "OPTIONS"
)

type ControllerRoute struct {
	Method method
	Path   string
	Handle func(w http.ResponseWriter, req *http.Request) error
}

type Controller interface {
	GetRoutes() []ControllerRoute
}
