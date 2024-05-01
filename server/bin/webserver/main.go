package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"pacstall.dev/webserver/config"
	"pacstall.dev/webserver/controllers"
	ps_api "pacstall.dev/webserver/controllers/pacscripts"
	repology_api "pacstall.dev/webserver/controllers/repology"
	urlshortener "pacstall.dev/webserver/controllers/url_shortener"
	"pacstall.dev/webserver/log"
	"pacstall.dev/webserver/model"
	grs "pacstall.dev/webserver/services/git_resolver_service"
	"pacstall.dev/webserver/services/matomo_tracker"
	pkgcache "pacstall.dev/webserver/services/package_cache"
	"pacstall.dev/webserver/services/parser"
	"pacstall.dev/webserver/services/repology"
	"pacstall.dev/webserver/services/server"
	ssr "pacstall.dev/webserver/services/serverside_render"
	"pacstall.dev/webserver/types/controller"
	"pacstall.dev/webserver/types/repository"
	"pacstall.dev/webserver/types/service"
)

func printLogo(conf config.ServerConfiguration) {
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
		built at: ` + conf.BuildDate + `
		version: ` + conf.Version + `
		mode: ` + func() string {
		if conf.Production {
			return "production"
		}

		return "development"
	}() + `
`))
}

func main() {
	// Load configuration
	globalConfiguration := config.Parse()

	if globalConfiguration.ServerConfiguration.Production {
		log.SetLogLevel(log.Level.Info)
	} else {
		log.SetLogLevel(log.Level.Debug)
	}

	startedAt := time.Now()

	// Initialize repositories
	databaseConnection := model.Connect(globalConfiguration.DatabaseConfiguration)
	var repologyProjectProviderRepository repository.RepologyProjectProviderRepository = model.InitRepologyProjectProviderRepository(databaseConnection)
	var repologyProjectRepository repository.RepologyProjectRepository = model.InitRepologyProjectRepository(databaseConnection)
	var shortenedLinkRepository repository.ShortenedLinkRepository = model.InitShortenedLinkRepository(databaseConnection)

	// Initialize services
	var repologyService service.RepologyService = repology.New(
		globalConfiguration.RepologyConfiguration,
		repologyProjectRepository,
		repologyProjectProviderRepository,
	)
	var packageCacheService service.PackageCacheService = pkgcache.New()
	var serverService service.ServerService = server.New(
		globalConfiguration.ServerConfiguration,
		packageCacheService,
	)
	var parserService service.ParserService = parser.New(
		globalConfiguration.PacstallProgramsConfiguration,
		globalConfiguration.ServerConfiguration,
		globalConfiguration.RepologyConfiguration,
		repologyService,
		grs.New(grs.NewShellGitCommitResolver()),
		packageCacheService,
	)
	var ssrService service.ServerSideRenderService = ssr.New(packageCacheService)
	var matomoTrackerService service.MatomoTrackerService = matomo_tracker.New()

	// Initialize controllers
	var urlShortenerController controller.Controller = urlshortener.New(
		globalConfiguration.MatomoConfiguration,
		shortenedLinkRepository,
		matomoTrackerService,
	)
	var repologyController controller.Controller = repology_api.New(
		globalConfiguration.ServerConfiguration,
		packageCacheService,
	)
	var packageController controller.Controller = ps_api.New(
		globalConfiguration.ServerConfiguration,
		packageCacheService,
	)

	// Initialize ControllersManager
	controllesManager := controllers.New(
		serverService.Router(),
		[]controller.Controller{
			repologyController,
			packageController,
			urlShortenerController,
		},
	)

	// Startup
	printLogo(globalConfiguration.ServerConfiguration)

	ssrService.EnableServerSideRendering()
	controllesManager.RegisterRoutes()

	log.Info("attempting to start TCP listener")

	server.OnServerOnline(func() {
		log.NotifyCustom("🚀 Startup 🧑‍🚀", "successfully started up.")
		log.Info("server is now online on port %v.\n", globalConfiguration.ServerConfiguration.Port)

		log.Info("booted in %v\n", color.GreenString("%v", time.Since(startedAt)))

		parserService.ScheduleRefresh(globalConfiguration.PacstallProgramsConfiguration.UpdateInterval)
		log.Info("scheduled pacscripts to auto-refresh every %v", globalConfiguration.PacstallProgramsConfiguration.UpdateInterval)

		if globalConfiguration.RepologyConfiguration.Enabled {
			repologyService.ScheduleRefresh(globalConfiguration.RepologyConfiguration.UpdateInterval)
			log.Info("scheduled repology to auto-refresh every %v", globalConfiguration.RepologyConfiguration.UpdateInterval)
		} else {
			log.Warn("repository repology is disabled")
		}
	})

	serverService.Listen()
}
