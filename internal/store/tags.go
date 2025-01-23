package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/karbasia/karbasi.dev/internal/database"
)

type TagCore struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type Tag struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type TagStore struct {
	db *sql.DB
}

func (s *TagStore) Create(ctx context.Context, tag *Tag) error {
	query := `
		INSERT INTO TAGS(name)
		VALUES($1)
		RETURNING id, created_at, updated_at
	`
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		tag.Name,
	).Scan(
		&tag.ID,
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *TagStore) GetAll(ctx context.Context) ([]Tag, error) {
	query := `
		SELECT id, name, created_at, updated_at
		FROM tags
		WHERE deleted_at IS NULL
		ORDER BY name
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []Tag{}
	for rows.Next() {
		t := Tag{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.CreatedAt,
			&t.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return tags, nil
}
