package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"pacstall.dev/website/cfg"
	"pacstall.dev/website/pacpkgs"
	"pacstall.dev/website/serverlib"
)

func setupRequests() {
	serverlib.HandleRequest("/api/packages", serverlib.HttpsMethods.GET, func(w http.ResponseWriter, req *http.Request) {
		packages, err := pacpkgs.GetPackages()
		if err != nil {
			log.Printf("Error > Could not parse packagelist file. Setting response 500.\n%v\n", err)
			w.WriteHeader(500)
			return
		}

		json, err := json.Marshal(packages)
		if err != nil {
			log.Printf("Error > Could not marshal to json. Setting response 500.\n%v\n", err)
			w.WriteHeader(500)
		}

		w.Write(json)
	})
}

func main() {
	port := cfg.Config.TCPServer.Port

	log.Println("Attempting to register http requests")
	setupRequests()
	log.Println("Successfully registered http requests")

	log.Println("Attempting to parse existing packages")
	pacpkgs.LoadPackages()
	log.Println("Successfully parsed existing packages")

	log.Println("Attempting to schedule package auto-refresh")
	pacpkgs.ScheduleRefresh(time.Duration(cfg.Config.PacstallPrograms.UpdateInterval) * time.Millisecond)
	log.Println("Successfully scheduled package auto-refresh")

	log.Println("Attempting to start TCP listener")

	serverlib.OnServerOnline(func() {
		log.Printf("Server is now online on port %v.\n", port)
	})

	serverlib.Serve(port)
}
