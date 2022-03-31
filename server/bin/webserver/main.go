package main

import (
	"fmt"
	"log"
	"time"

	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/featureflag"
	"pacstall.dev/webserver/listener"
	"pacstall.dev/webserver/pacscript"
	"pacstall.dev/webserver/pacscript/pshttphandle"
	pacssr "pacstall.dev/webserver/pacscript/ssr"
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
	router := listener.Router()

	/* Packages */
	pacssr.EnableSSR()

	router.HandleFunc("/api/packages", pshttphandle.GetPacscriptListHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}", pshttphandle.GetPacscriptHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/requiredBy", pshttphandle.GetPacscriptRequiredByHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/dependencies", pshttphandle.GetPacscriptDependenciesHandle).Methods("GET")

	/* Feature Flags */
	router.HandleFunc("/api/feature-flags", featureflag.GetFeatureFlags).Methods("GET")
}

func main() {
	startedAt := time.Now()
	port := config.Config.TCPServer.Port
	refreshTimer := time.Duration(config.Config.PacstallPrograms.UpdateInterval) * time.Millisecond

	setupRequests()
	log.Println("Registered http requests")

	log.Println("Attempting to start TCP listener")

	listener.OnServerOnline(func() {
		log.Printf("Server is now online on port %v.\n", port)

		log.Println("Attempting to parse existing pacscripts")
		pacscript.Load()
		pacscript.ScheduleRefresh(refreshTimer)
		log.Println("Scheduled pacscripts to auto-refresh every", refreshTimer)

		printLogo()
		log.Printf("Booted in %v%v%v\n", "\033[32m", time.Since(startedAt), "\033[0m")
	})

	listener.Listen(port)
}
