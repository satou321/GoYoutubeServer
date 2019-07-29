package main

import (
	"github.com/golang/go/src/strconv"
	"github.com/skratchdot/open-golang/open"
	"log"

	"github.com/GoYoutubeServer/server/config"
	"github.com/GoYoutubeServer/server/controllers"
)

func init() {
	port := strconv.Itoa(config.Config.Port)
	log.Print(open.Run("http://localhost:" + port + "/r"))
}

func main() {
	//var cfg Config
	cfg := config.Config

	err := controllers.StartWebServer(cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
