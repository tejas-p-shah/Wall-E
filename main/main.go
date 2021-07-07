package main

import (
	"log"

	"github.com/tejas-p-shah/Wall-E/bootstrap"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	bootstrap.BootApplication()
}
