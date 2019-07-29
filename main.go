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
	print(port)

	//http.ListenAndServe(":"+port, Log(router))
	log.Fatal(http.ListenAndServe(":"+port, nil))

	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = "80"
	//}
	//err := controllers.StartWebServer(port)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
