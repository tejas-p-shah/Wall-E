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

// func GetSecretKey() string {
// 	secretKey, exists := os.LookupEnv("SECRETKEY")
// 	if !exists {
// 		log.Fatal("Could not Found SECRETKEY in .env file")
// 	}

// 	return secretKey
// }
