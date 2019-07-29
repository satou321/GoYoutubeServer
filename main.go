package main

import (
	"fmt"
	"github.com/satou321/GoYoutubeServer/server/controllers"
	"log"
	"os"
	"strconv"
)

func main() {
	port := getPort()
	err := controllers.StartWebServer(port)
	if err != nil {
		log.Fatal(err)
	}
}
func getPort() int {
	p, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	if p == 0 {
		return 80
	}
	return p
}
