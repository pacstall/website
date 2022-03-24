package main

import (
	"fmt"
	"log"
	"time"

	"pacstall.dev/website/config"
	"pacstall.dev/website/featureflag"
	"pacstall.dev/website/listener"
	"pacstall.dev/website/pacscript"
	"pacstall.dev/website/pacscript/pshttphandle"
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
	router.HandleFunc("/api/packages", pshttphandle.GetPackageListHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}", pshttphandle.GetPackageHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/requiredBy", pshttphandle.GetPackagesRequiredByHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/dependencies", pshttphandle.GetPackageDependenciesHandle).Methods("GET")

	/* Feature Flags */
	router.HandleFunc("/api/feature-flags", featureflag.GetFeatureFlags).Methods("GET")
}

func main() {
	startedAt := time.Now()
	port := config.Config.TCPServer.Port

	setupRequests()
	log.Println("Successfully registered http requests")

	log.Println("Attempting to start TCP listener")

	listener.OnServerOnline(func() {
		log.Printf("Server is now online on port %v.\n", port)

		log.Println("Attempting to parse existing packages")
		pacscript.LoadPackages()
		pacscript.ScheduleRefresh(time.Duration(config.Config.PacstallPrograms.UpdateInterval) * time.Millisecond)
		log.Println("Successfully scheduled package auto-refresh")

		printLogo()
		log.Printf("Booted in %v%v%v\n", "\033[32m", time.Since(startedAt), "\033[0m")
	})

	listener.Listen(port)
}
