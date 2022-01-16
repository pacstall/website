package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"pacstall.dev/website/cfg"
	"pacstall.dev/website/pacpkgs"
	"pacstall.dev/website/serverlib"
)

func printLogo() {
	fmt.Println("\033[34m" + `
     ____                  __        ____         
    / __ \____ ___________/ /_____ _/ / /         
   / /_/ / __ '/ ___/ ___/ __/ __ '/ / /          
  / ____/ /_/ / /__(__  ) /_/ /_/ / / /           
 /_/   _\__,_/\___/____/\__/\__,_/_/_/            
| |     / /__  / /_ / ___/___  ______   _____  _____
| | /| / / _ \/ __ \\__ \/ _ \/ ___/ | / / _ \/ ___/
| |/ |/ /  __/ /_/ /__/ /  __/ /   | |/ /  __/ /    
|__/|__/\___/_.___/____/\___/_/    |___/\___/_/     
         coded by saenai255, owned by Pacstall Org		  
	` + "\033[0m")
}

func setupRequests() {
	serverlib.HandleRequest("/api/packages", serverlib.HttpsMethods.GET, func(w http.ResponseWriter, req *http.Request) {
		packages := pacpkgs.GetPackages()

		json, err := json.Marshal(packages)
		if err != nil {
			log.Printf("Could not marshal to json. Setting response 500.\n%v\n", err)
			w.WriteHeader(500)
		}

		w.Write(json)
	})
}

func main() {
	startedAt := time.Now()
	port := cfg.Config.TCPServer.Port

	setupRequests()
	log.Println("Successfully registered http requests")
	log.Println("Attempting to parse existing packages")
	pacpkgs.LoadPackages()

	pacpkgs.ScheduleRefresh(time.Duration(cfg.Config.PacstallPrograms.UpdateInterval) * time.Millisecond)
	log.Println("Successfully scheduled package auto-refresh")
	log.Println("Attempting to start TCP listener")

	serverlib.OnServerOnline(func() {
		log.Printf("Server is now online on port %v.\n", port)

		printLogo()
		log.Printf("Booted in %v%v%v\n", "\033[32m", time.Since(startedAt), "\033[0m")
	})

	serverlib.Serve(port)
}
