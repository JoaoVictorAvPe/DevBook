package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println("Port:", config.Port)
	fmt.Println("Stringconnection:", config.StringConnection)

	fmt.Print("Running API...")

	r := router.Create()

	log.Fatal(http.ListenAndServe(":" + config.Port, r))
}
