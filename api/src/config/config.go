package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var StringConnection, Port string

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port = os.Getenv("API_PORT")
	StringConnection = fmt.Sprintf("%s:%s@tcp(192.168.0.20:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
}
