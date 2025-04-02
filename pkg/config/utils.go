package config

import (
	"log"
	"os"
	"time"
)

func getEnv(key string) string {
	value, exits := os.LookupEnv(key)
	if value == "" || !exits {
		log.Fatalf("Error: Missing environment variable %s", key)
	}
	return value
}

func getDurationEnv(key string) time.Duration {
	timeout, err := time.ParseDuration(getEnv(key))
	if err != nil {
		log.Fatalf("Error: Can not parse %s env as duration", key)
	}
	return timeout
}
