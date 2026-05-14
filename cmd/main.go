package main

import (
	"fmt"
	"log"
	"net/http"
	"pjweb/internal/api"
	"pjweb/internal/config"
	"pjweb/internal/db"
)

func main() {
	applicationConfig, err := config.LoadConfig("../config.yaml")
	if err != nil {
		fmt.Print("config error\n")
		return
	}

	db, err := db.New(&applicationConfig.Database)
	if err != nil {
		fmt.Print("db init error\n")
		fmt.Print(err)
		return
	}

	routerService, err := api.NewRouterService(applicationConfig)
	if err != nil {
		fmt.Print("service error\n")
		return
	}
	router := api.NewMux(routerService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(server.ListenAndServe())

	db.Conn.Close()
}
