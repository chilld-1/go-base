package main

import (
	"xiangmu/config"
	"xiangmu/route"
)

func main() {
	config.InitConfig()
	r := route.SetupRouter()
	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8080"
	}

	r.Run(port)
}
