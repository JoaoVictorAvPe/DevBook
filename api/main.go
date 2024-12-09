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

	fmt.Printf("Running API on port %s", config.Port)

	r := router.Create()

	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
