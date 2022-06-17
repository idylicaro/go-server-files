package main

import (
	"log"

	routers "github.com/idylicaro/go-server-files/routers"
)

func main() {
	router := routers.SetupRouter()
	err := router.Run(":8080")
	if err != nil {
		// TODO: handle error properly
		log.Fatal(err)
	}
}
