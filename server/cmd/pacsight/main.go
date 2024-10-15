package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/fatih/color"
	"pacstall.dev/webserver/internal/pacsight/config"
	"pacstall.dev/webserver/internal/pacsight/repology"
	"pacstall.dev/webserver/internal/pacsight/rpcall"
	"pacstall.dev/webserver/pkg/common/log"
)

func printLogo() {
	logoColor := color.New(color.FgHiMagenta, color.Bold).SprintFunc()
	fmt.Println(logoColor(`
                                    d8b          888      888   
                                    Y8P          888      888   
                                                 888      888   
88888b.   8888b.   .d8888b .d8888b  888  .d88b.  88888b.  888888
888 "88b     "88b d88P"    88K      888 d88P"88b 888 "88b 888   
888  888 .d888888 888      "Y8888b. 888 888  888 888  888 888   
888 d88P 888  888 Y88b.         X88 888 Y88b 888 888  888 Y88b. 
88888P"  "Y888888  "Y8888P  88888P' 888  "Y88888 888  888  "Y888
888                                          888                
888                                     Y8b d88P                
888                                      "Y88P"                 

	coded by saenai255, owned by Pacstall Org		  
   `))
}

func main() {
	config.Init()
	printLogo()

	log.Info("booting up pacsight")

	go repology.ScheduleRefresh(config.Repology.RepologyUpdateInterval)
	rpcall.RegisterService()

	port := fmt.Sprintf(":%v", config.PacSight.Port)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to start server: %+v", err)
	}

	log.Info("server started on port " + port)

	if err := http.Serve(listener, nil); err != nil {
		log.Fatal("failed to serve: %+v", err)
	}

}
