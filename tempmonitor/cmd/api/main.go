package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"tempMonitor/internal/data"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
	"github.com/rabbitmq/amqp091-go"
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
	rabbitmq struct {
		url string
	}
}
type application struct {
	cfg      config
	logger   *slog.Logger
	db       *sql.DB
	rabbitmq *amqp091.Connection
	models   data.Models
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
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://dbuser:dbpwd@localhost:5431/tempmonitor?sslmode=disable", "PostgreSQL DSN")
	flag.StringVar(&cfg.rabbitmq.url, "rabbitmq-url", "amqp://guest:guest@localhost:5672/", "RabbitMQ URL")

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

func openRabbitMQ(cfg *config) (*amqp091.Connection, error) {
	conn, err := amqp091.Dial(cfg.rabbitmq.url)
	if err != nil {
		return nil, err
	}
	return conn, nil

}

func main() {
	cfg := newConfig()
	logger := newLogger()
	db, err := openDB(cfg)
	if err != nil {
		logger.Error("Cannot connect to DB", "err", err)
		os.Exit(exitFailure)
	}
	defer db.Close()

	rabbitmq, err := openRabbitMQ(cfg)
	if err != nil {
		logger.Error("Cannot connect to RabbitMQ", "err", err)
		os.Exit(exitFailure)
	}
	defer rabbitmq.Close()

	app := &application{
		cfg:      *cfg,
		logger:   logger,
		db:       db,
		rabbitmq: rabbitmq,
		models:   data.NewModels(db),
	}

	go app.listenToRabbitMQ("temperature_measurements")

	err = app.serve()

	if err != nil {
		logger.Error("Jebudut", "err", err)
		os.Exit(exitFailure)
	}

}
