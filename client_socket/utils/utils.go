package utils

import (
	"log"
	"os"
)

func Getenv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	return defaultValue
}

func Shutdown() {
	log.Println("Shuting down...")
	os.Exit(0)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
