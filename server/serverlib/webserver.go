package serverlib

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Serve(port int) {
	registerHealthCheck()
	for path, httpHandles := range handlers {
		http.HandleFunc(path, func(rw http.ResponseWriter, r *http.Request) {
			for _, httpHandle := range httpHandles {
				if strings.ToUpper(r.Method) == httpHandle.method {
					httpHandle.handler(rw, r)
					return
				}
			}

			rw.WriteHeader(405 /* Err Method Not Allowed */)
		})
	}

	go triggerServerOnline(port)
	sPort := fmt.Sprintf(":%v", port)
	err := http.ListenAndServe(sPort, nil)
	log.Panicf("Could not start TCP listener on port %v. Got error: %v", port, err)
}
