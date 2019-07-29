package main

import (
	"fmt"
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
	"strconv"
)

func main() {
	//var cfg Config
	//cfg := config.Config
	port, _ := strconv.Atoi(os.Args[1])
	fmt.Println(port)
	//port := (os.Getenv("$PORT")||"8083")
	if port == 0 {
		log.Fatal("PORT must be set", port)
	}

	err := controllers.StartWebServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
