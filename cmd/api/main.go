package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/karbasia/karbasi.dev/internal/database"
	"github.com/karbasia/karbasi.dev/internal/store"
	"github.com/karbasia/karbasi.dev/internal/version"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

type config struct {
	baseURL  string
	httpPort int
	db       struct {
		dsn         string
		automigrate bool
	}
	jwt struct {
		accessSecretKey  string
		refreshSecretKey string
	}
}

type application struct {
	config config
	store  store.Storage
	logger *slog.Logger
	wg     sync.WaitGroup
}

func run(logger *slog.Logger) error {
	var cfg config

	flag.StringVar(&cfg.baseURL, "base-url", "http://localhost:8080", "base URL for the application")
	flag.IntVar(&cfg.httpPort, "http-port", 8080, "port to listen on for HTTP requests")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "db.sqlite", "sqlite3 DSN")
	flag.BoolVar(&cfg.db.automigrate, "db-automigrate", true, "run migrations on startup")
	flag.StringVar(&cfg.jwt.accessSecretKey, "jwt-access-secret-key", "vyiwjr425wpr277oxf34tcmg73mmkcks", "secret key for access JWT")
	flag.StringVar(&cfg.jwt.refreshSecretKey, "jwt-refresh-secret-key", "tqAp56ce2i2XmpoAsubxkgV0ThBYYFBV", "secret key for refresh JWT")

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		return err
	}
	defer db.Close()

	store := store.New(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	return app.serveHTTP()
}
