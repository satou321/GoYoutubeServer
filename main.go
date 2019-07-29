package main

import (
	"github.com/satou321/GoYoutubeServer/server/config"
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
)

func main() {
	//var cfg Config
	cfg := config.Config

	err := controllers.StartWebServer(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
