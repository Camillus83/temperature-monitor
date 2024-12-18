package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}
type application struct {
	cfg    config
	logger *slog.Logger
	db     *sql.DB
}

func newLogger() *slog.Logger {
	return slog.New(
		log.NewWithOptions(
			os.Stdout,
			log.Options{
				ReportCaller:    true,
				ReportTimestamp: true,
				TimeFormat:      time.DateTime,
				Prefix:          "TempMonitor 🌡️",
			},
		),
	)
}

func newConfig() *config {
	cfg := &config{}
	flag.IntVar(&cfg.port, "port", 4000, "WEBAPP server port")
	flag.StringVar(
		&cfg.env,
		"env",
		"development",
		"Environment (development|staging|production)",
	)
	flag.StringVar(&cfg.db.dsn, "dsn", "tempmonitor.db", "SQLite DSN")

	flag.Parse()
	return cfg
}

func main() {
	cfg := newConfig()
	logger := newLogger()

	db, err := sql.Open("sqlite3", cfg.db.dsn)
	if err != nil {
		logger.Error("Failed to open SQLite db", err)
	}
	defer db.Close()

	app := &application{
		cfg:    *cfg,
		logger: logger,
		db:     db,
	}

	fmt.Println(app.cfg.port)
}
