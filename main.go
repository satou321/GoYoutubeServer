package main

import (
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
)

func main() {
	port := getPort()
	err := controllers.StartWebServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return p
	}
	return "80"
}
