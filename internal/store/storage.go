package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetAllByTag(context.Context, string) ([]Post, error)
		GetBySlug(context.Context, string) (*Post, bool, error)
		GetAll(context.Context) ([]Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
		GetByID(context.Context, int) (*User, bool, error)
		GetByEmail(context.Context, string) (*User, bool, error)
		UpdateHashedPassword(context.Context, int, string) error
		GetAll(context.Context) ([]User, error)
	}
	Tags interface {
		Create(context.Context, *Tag) error
		GetAll(context.Context) ([]Tag, error)
	}
}

func New(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
		Tags:  &TagStore{db},
	}
}
