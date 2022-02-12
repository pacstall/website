package main

import (
	"fmt"
	"log"
	"time"

	"pacstall.dev/website/cfg"
	"pacstall.dev/website/fflags"
	"pacstall.dev/website/pacpkgs"
	"pacstall.dev/website/svlib"
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
	router := svlib.Router()

	/* Packages */
	router.HandleFunc("/api/packages", pacpkgs.GetPackageListHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}", pacpkgs.GetPackageHandle).Methods("GET")

	/* Feature Flags */
	router.HandleFunc("/api/feature-flags", fflags.GetFeatureFlags).Methods("GET")
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

	svlib.OnServerOnline(func() {
		log.Printf("Server is now online on port %v.\n", port)

		printLogo()
		log.Printf("Booted in %v%v%v\n", "\033[32m", time.Since(startedAt), "\033[0m")
	})

	svlib.Serve(port)
}
