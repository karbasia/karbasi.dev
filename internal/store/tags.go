package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/karbasia/karbasi.dev/internal/database"
)

type Tag struct {
	ID        int        `json:"id,omitzero,omitempty" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt *time.Time `json:"created_at,omitzero" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitzero" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitzero" db:"deleted_at"`
	PostCount *int       `json:"post_count,omitempty"`
}

type TagStore struct {
	db *sql.DB
}

func (s *TagStore) Create(ctx context.Context, tag *Tag) error {
	query := `
		INSERT INTO tags(name)
		VALUES($1)
		ON CONFLICT(name) DO UPDATE SET deleted_at = NULL
		RETURNING id, created_at, updated_at;
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

func (s *TagStore) Update(ctx context.Context, tag *Tag) error {
	query := `
		UPDATE tags SET (name, deleted_at) =
		($1, $2)
		WHERE id = $3
		RETURNING created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		tag.Name,
		tag.DeletedAt,
		tag.ID,
	).Scan(
		&tag.CreatedAt,
		&tag.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *TagStore) GetAll(ctx context.Context, showDeleted bool) ([]Tag, error) {
	filterParam := ""
	if !showDeleted {
		filterParam = "WHERE deleted_at IS NULL"
	}
	query := fmt.Sprintf(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM tags
		%s
		ORDER BY name
	`, filterParam)

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
			&t.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return tags, nil
}

func (s *TagStore) GetAllByPostCount(ctx context.Context) ([]Tag, error) {
	query := `
		SELECT t.id, t.name, t.created_at, t.updated_at, t.deleted_at, COUNT(p.id) AS post_count
		FROM tags t
		LEFT JOIN posts_to_tags pt ON t.id = pt.tag_id
		LEFT JOIN posts p ON p.id = pt.post_id
		WHERE p.deleted_at IS NULL 
			AND t.deleted_at IS NULL
		GROUP BY t.id
		ORDER BY COUNT(p.id) DESC
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
			&t.DeletedAt,
			&t.PostCount,
		)
		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return tags, nil
}
