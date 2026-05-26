package main

import (
	"fmt"
	"pjweb/internal/app"
	"pjweb/internal/config"
)

func main() {
	config, err := config.LoadConfig("../config.test.yaml")
	if err != nil {
		fmt.Printf("config error\n")
		return
	}

	app, err := app.New(config)

	err = app.Run()
	if err != nil {
		fmt.Printf("app fun failed")
	}
}
