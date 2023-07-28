package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/repology"
	"pacstall.dev/webserver/server"
	ps_api "pacstall.dev/webserver/server/api/pacscripts"
	repology_api "pacstall.dev/webserver/server/api/repology"
	urlshortener "pacstall.dev/webserver/server/api/url_shortener"
	pac_ssr "pacstall.dev/webserver/server/ssr/pacscript"
	"pacstall.dev/webserver/types/pac/parser"
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
	pac_ssr.EnableSSR()
	router.HandleFunc("/api/repology", repology_api.GetRepologyPackageListHandle).Methods("GET")
	router.HandleFunc("/api/packages", ps_api.GetPacscriptListHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}", ps_api.GetPacscriptHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/requiredBy", ps_api.GetPacscriptRequiredByHandle).Methods("GET")
	router.HandleFunc("/api/packages/{name}/dependencies", ps_api.GetPacscriptDependenciesHandle).Methods("GET")

	/* Shortened Links - Must be last as it functions as a catch-all trap */
	router.HandleFunc("/q/{linkId}", urlshortener.GetShortenedLinkRedirectHandle).Methods("GET")
}

func main() {
	if config.Production {
		log.SetLogLevel(log.Level.Info)
	} else {
		log.SetLogLevel(log.Level.Debug)
	}

	startedAt := time.Now()

	printLogo()

	setupRequests()
	log.Info("Registered http requests")

	log.Info("Attempting to start TCP listener")

	server.OnServerOnline(func() {
		log.NotifyCustom("🚀 Startup 🧑‍🚀", "Successfully started up.")
		log.Info("Server is now online on port %v.\n", config.Port)

		log.Info("Booted in %v\n", color.GreenString("%v", time.Since(startedAt)))

		parser.ScheduleRefresh(config.UpdateInterval)
		log.Info("Scheduled pacscripts to auto-refresh every %v", config.UpdateInterval)

		if config.Repology.Enabled {
			repology.ScheduleRefresh(config.RepologyUpdateInterval)
			log.Info("Scheduled repology to auto-refresh every %v", config.RepologyUpdateInterval)
		} else {
			log.Warn("Repology is disabled. Pacstall will not be able to fetch package data from Repology.")
		}
	})

	server.Listen(config.Port)
}
