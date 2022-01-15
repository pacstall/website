package serverlib

import "net/http"

type httpMethod string

type httpMethods struct {
	GET    httpMethod
	POST   httpMethod
	PUT    httpMethod
	DELETE httpMethod
}

var HttpsMethods httpMethods = httpMethods{
	GET:    httpMethod("GET"),
	POST:   httpMethod("POST"),
	PUT:    httpMethod("PUT"),
	DELETE: httpMethod("DELETE"),
}

type httpHandler struct {
	method  string
	handler func(w http.ResponseWriter, req *http.Request)
}

var handlers = make(map[string][]httpHandler)

func HandleRequest(pathOrPattern string, method httpMethod, handler func(w http.ResponseWriter, req *http.Request)) {
	_, ok := handlers[pathOrPattern]
	if !ok {
		handlers[pathOrPattern] = make([]httpHandler, 1)
	}

	handlers[pathOrPattern] = append(handlers[pathOrPattern], httpHandler{
		method:  string(method),
		handler: handler,
	})
}
