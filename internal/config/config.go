package config

import "os"

type Config struct {
	ServerPort   string
	DatabasePath string
}

func Load() Config {
	cfg := Config{
		ServerPort:   getEnv("SERVER_PORT", ":8080"),
		DatabasePath: getEnv("DATABASE_PATH", "movies.db"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}
