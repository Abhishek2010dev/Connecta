package config

import (
	"log"
	"os"
)

func getEnv(key string) string {
	value, exits := os.LookupEnv(key)
	if value == "" || !exits {
		log.Fatalf("Error: Missing environment variable %s", key)
	}
	return value
}
