package main

import (
	"fmt"
	"log"
	"net/http"
	"pjweb/internal/api"
	"pjweb/internal/config"
)

func main() {
	applicationConfig, err := config.LoadConfig("../config.yaml")
	if err != nil {
		fmt.Print("config error")
		return
	}

	routerService, err := api.RouterServiceByConfig(applicationConfig)
	if err != nil {
		fmt.Print("service error")
	}
	router := api.NewRouter(routerService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())
}
