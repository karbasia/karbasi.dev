package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/karbasia/karbasi.dev/assets"
	"github.com/pressly/goose/v3"

	_ "modernc.org/sqlite"
)

const DefaultTimeout = 30 * time.Second

func New(dsn string, automigrate bool) (*sql.DB, error) {

	pragmas := "?_fk=1"
	db, err := sql.Open("sqlite", dsn+pragmas)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Hour)

	if automigrate {
		goose.SetBaseFS(assets.EmbeddedFiles)

		if err := goose.SetDialect("sqlite"); err != nil {
			return nil, err
		}
		if err := goose.Up(db, "migrations"); err != nil {
			return nil, err
		}
	}

	return db, nil
}
