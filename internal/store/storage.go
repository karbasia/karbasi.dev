package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		Update(context.Context, *Post) error
		GetAllByTag(context.Context, string) ([]Post, error)
		GetBySlug(context.Context, string) (*Post, bool, error)
		GetAll(context.Context, bool) ([]Post, error)
	}
	Users interface {
		Create(context.Context, *User) error
		Update(context.Context, *User) error
		GetByID(context.Context, int) (*User, bool, error)
		GetByEmail(context.Context, string) (*User, bool, error)
		GetAll(context.Context) ([]User, error)
	}
	Tags interface {
		Create(context.Context, *Tag) error
		Update(context.Context, *Tag) error
		GetAll(context.Context, bool) ([]Tag, error)
		GetAllByPostCount(context.Context) ([]Tag, error)
	}
	Files interface {
		Create(context.Context, *File) error
		GetAll(context.Context) ([]File, error)
		GetByID(context.Context, int) (*File, bool, error)
		GetByName(context.Context, string) (*File, bool, error)
	}
}

func New(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
		Users: &UserStore{db},
		Tags:  &TagStore{db},
		Files: &FileStore{db},
	}
}
