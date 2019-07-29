package main

import (
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
)

func main() {
	//var cfg Config
	//cfg := config.Config
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	err := controllers.StartWebServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
