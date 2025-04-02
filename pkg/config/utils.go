package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

func LoadEnv(key string) string {
	value, exits := os.LookupEnv(key)
	if value == "" || !exits {
		log.Fatalf("Error: Missing environment variable %s", key)
	}
	return value
}

func LoadEnvDuration(key string) time.Duration {
	timeout, err := time.ParseDuration(LoadEnv(key))
	if err != nil {
		log.Fatalf("Error: Can not parse %s env as duration", key)
	}
	return timeout
}

func LoadEnvInt(key string) int {
	value, err := strconv.Atoi(LoadEnv(key))
	if err != nil {
		log.Fatalf("Error: Can not parse %s env as int", key)
	}
	return value
}
