package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Running API...")

	r := router.Create()

	log.Fatal(http.ListenAndServe(":5000", r))
}
