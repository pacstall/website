package main

import (
	"fmt"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/ztrue/shutdown"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/server"
	ps_api "pacstall.dev/webserver/server/api/pacscripts"
	repology_api "pacstall.dev/webserver/server/api/repology"
	pac_ssr "pacstall.dev/webserver/server/ssr/pacscript"
	conf_type "pacstall.dev/webserver/types/config"
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
}

func main() {
	// Use dummy log config until the real one is parsed
	log.Init(conf_type.LoggingConfig{
		DiscordToken:     "",
		DiscordChannelID: "",
		DiscordTags:      "",
		DiscordEnabled:   false,
	})

	config.Load()
	log.Init(config.Logging)

	shutdown.Add(func() {
		log.Notify("Shutting down...")
	})

	go shutdown.Listen(syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	startedAt := time.Now()
	port := config.TCPServer.Port
	refreshTimer := config.PacstallPrograms.UpdateInterval

	setupRequests()
	log.Info("Registered http requests")

	log.Info("Attempting to start TCP listener")

	server.OnServerOnline(func() {
		log.NotifyCustom("🚀 Startup 🧑‍🚀", "Successfully started up.")
		log.Info("Server is now online on port %v.\n", port)

		printLogo()
		log.Info("Booted in %v\n", color.GreenString("%v", time.Since(startedAt)))

		log.Info("Attempting to parse existing pacscripts")
		parser.ParseAll()
		parser.ScheduleRefresh(refreshTimer)
		log.Info("Scheduled pacscripts to auto-refresh every", refreshTimer)
	})

	server.Listen(port)
}
