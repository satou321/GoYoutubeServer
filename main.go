package main

import (
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
)

func main() {
	//port := getPort()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	err := controllers.StartWebServer(port)
	if err != nil {
		log.Fatal(err)
	}
}

//func getPort() int {
//p, err := strconv.Atoi(os.Getenv("PORT"))
//if err != nil {
//	fmt.Println(err)
//}
//fmt.Println(p)
//if p == 0 {
//	return 80
//}
//return p
//}
