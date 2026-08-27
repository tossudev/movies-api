package config

import "os"

type Config struct {
	Port     string
	Database string
}

func Load() Config {
	cfg := Config{
		Port:     getEnv("PORT", "8080"),
		Database: getEnv("DATABASE", "./movies.db"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
