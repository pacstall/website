package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	psapi "pacstall.dev/webserver/server/api/pacscripts"
	repology_api "pacstall.dev/webserver/server/api/repology"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/types/pac/parser"
	"pacstall.dev/webserver/server"
	pacssr "pacstall.dev/webserver/server/ssr/pacscript"
)

func printLogo() {
	logoColor := color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	fmt.Println(logoColor(`
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
   `))
}

func setupRequests() {
	router := server.Router()

	/* Packages */
	pacssr.EnableSSR()

	router.HandleFunc("/api/repology", repology_api.GetRepologyPackageListHandle).Methods("GET")
	router.HandleFunc("/api/packages", psapi.GetPacscriptListHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}", psapi.GetPacscriptHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/requiredBy", psapi.GetPacscriptRequiredByHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/dependencies", psapi.GetPacscriptDependenciesHandle).Methods("GET")
}

func main() {
	config.Load()
	log.Init(config.Logging.FancyLogs, config.Logging.Level)

	startedAt := time.Now()
	port := config.TCPServer.Port
	refreshTimer := config.PacstallPrograms.UpdateInterval

	setupRequests()
	log.Info.Println("Registered http requests")

	log.Info.Println("Attempting to start TCP listener")

	server.OnServerOnline(func() {
		log.Info.Printf("Server is now online on port %v.\n", port)

		printLogo()
		log.Info.Printf("Booted in %v\n", color.GreenString("%v", time.Since(startedAt)))

		log.Info.Println("Attempting to parse existing pacscripts")
		parser.ParseAll()
		parser.ScheduleRefresh(refreshTimer)
		log.Info.Println("Scheduled pacscripts to auto-refresh every", refreshTimer)
	})

	server.Listen(port)
}
