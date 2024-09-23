package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"pacstall.dev/webserver/internal/pacnexus/config"
	"pacstall.dev/webserver/internal/pacnexus/server"
	ps_api "pacstall.dev/webserver/internal/pacnexus/server/api/pacscripts"
	repology_api "pacstall.dev/webserver/internal/pacnexus/server/api/repology"
	urlshortener "pacstall.dev/webserver/internal/pacnexus/server/api/url_shortener"
	pac_ssr "pacstall.dev/webserver/internal/pacnexus/server/ssr/pacscript"
	"pacstall.dev/webserver/internal/pacnexus/types/pac/parser"
	globalConfig "pacstall.dev/webserver/pkg/common/config"
	"pacstall.dev/webserver/pkg/common/log"
	"pacstall.dev/webserver/pkg/common/pacsight"
)

func printLogo() {
	logoColor := color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	fmt.Println(logoColor(`
88888b.   8888b.   .d8888b 88888b.   .d88b.  888  888 888  888 .d8888b 
888 "88b     "88b d88P"    888 "88b d8P  Y8b  Y8bd8P' 888  888 88K     
888  888 .d888888 888      888  888 88888888   X88K   888  888 "Y8888b.
888 d88P 888  888 Y88b.    888  888 Y8b.     .d8""8b. Y88b 888      X88
88888P"  "Y888888  "Y8888P 888  888  "Y8888  888  888  "Y88888  88888P'
888                                                                    
888                                                                    
888                                                                    
	
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
	config.Init()

	if globalConfig.Production {
		log.SetLogLevel(log.Level.Info)
	} else {
		log.SetLogLevel(log.Level.Debug)
	}

	startedAt := time.Now()

	printLogo()

	setupRequests()
	log.Info("registered http requests")

	pacsightRpc, err := pacsight.NewPacsightRpcService("localhost", 8080)
	if err != nil {
		log.Error("failed to create pacsight rpc service: %+v", err)
		return
	}
	log.Info("connected to pacsight rpc service")

	log.Info("attempting to start tcp listener")

	server.OnServerOnline(func() {
		log.Info("server is now online on port %v.\n", config.PacNexus.Port)

		log.Info("booted in %v\n", color.GreenString("%v", time.Since(startedAt)))

		parser.ScheduleRefresh(config.PacstallPrograms.UpdateInterval, pacsightRpc)
		log.Info("scheduled pacscripts to auto-refresh every %v", config.PacstallPrograms.UpdateInterval)
	})

	server.Listen(config.PacNexus.Port)
}
