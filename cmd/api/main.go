package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/karbasia/karbasi.dev/internal/database"
	"github.com/karbasia/karbasi.dev/internal/env"
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

	cfg.baseURL = env.GetString("BASE_URL", "http://localhost:8080")
	cfg.httpPort = env.GetInt("HTTP_PORT", 8080)
	cfg.db.dsn = env.GetString("DB_DSN", "/data/db.sqlite")
	cfg.db.automigrate = env.GetBool("DB_AUTOMIGRATE", true)
	cfg.jwt.accessSecretKey = env.GetString("JWT_ACCESS_SECRET", "vyiwjr425wpr277oxf34tcmg73mmkcks")
	cfg.jwt.refreshSecretKey = env.GetString("JWT_REFRESH_SECRET", "tqAp56ce2i2XmpoAsubxkgV0ThBYYFBV")

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
