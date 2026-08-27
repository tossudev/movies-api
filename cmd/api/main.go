package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"movies-api/internal/config"
	"movies-api/internal/db"
	"movies-api/internal/routes"
)

func main() {
	cfg := config.Load()
	database, err := db.Open(cfg.Database)
	if err != nil {
		slog.Error("opening database", "err", err)
		return
	}
	defer database.Close()

	if err := db.Migrate(database); err != nil {
		slog.Error("migrating database", "err", err)
		return
	}

	host, err := os.Hostname()
	if err != nil {
		slog.Error("fetching hostname", "err", err)
		host = "localhost"
	}
	slog.Info("starting http server", "port", cfg.Port, "url", fmt.Sprintf("http://%s:%s/", host, cfg.Port))

	slog.Error("starting server", "err", http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), routes.New(database)))
}
