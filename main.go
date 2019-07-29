package main

import (
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := controllers.StartWebServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
