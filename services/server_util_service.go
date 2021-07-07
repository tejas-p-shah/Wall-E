package services

import (
	"log"
	"os"
)

func GetPort() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Could not Found PORT in .env file")
	}

	return port
}
