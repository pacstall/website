package serverlib

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func registerHealthCheck() {
	HandleRequest("/api/health", HttpsMethods.GET, func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
	})
}

var onServerOnlineHandlers []func()

func OnServerOnline(handle func()) {
	onServerOnlineHandlers = append(onServerOnlineHandlers, handle)
}

func triggerServerOnline(port int) {
	retryDelay := 100
	timeout := 5000
	for ; timeout > 0; timeout -= retryDelay {
		// Retry contacting server until online
		time.Sleep(time.Millisecond * time.Duration(retryDelay))
		_, err := http.Get(fmt.Sprintf("http://localhost:%v", port))
		if err == nil {
			break
		}
	}

	if timeout <= 0 {
		log.Fatalln("TCP server bootstrapping timed out.")
		return
	}

	for _, handler := range onServerOnlineHandlers {
		handler()
	}
}
