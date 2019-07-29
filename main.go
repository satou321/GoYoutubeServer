package main

import (
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
)

func main() {
	//var cfg Config
	//cfg := config.Config
	//port, _ := strconv.Atoi(os.Args[1])
	port := getPort()
	//if port == 0 {
	//	log.Fatal("PORT must be set", port)
	//}

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
