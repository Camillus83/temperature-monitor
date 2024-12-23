package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
)

const (
	exitSuccess int = iota
	exitFailure
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
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://dbuser@dbpwd:localhost/tempmonitor?sslmode=false", "PostgreSQL DSN")

	flag.Parse()
	return cfg
}

func openDB(cfg *config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil

}

func main() {
	cfg := newConfig()
	logger := newLogger()
	db, err := openDB(cfg)
	if err != nil {
		logger.Error("Cannot connect to DB", "err", err)
		os.Exit(exitFailure)
	}

	app := &application{
		cfg:    *cfg,
		logger: logger,
		db:     db,
	}

	log.Info("hi")
	log.Info("Running on Port: ", app.cfg.port)

}
