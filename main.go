package main

import (
	//"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
