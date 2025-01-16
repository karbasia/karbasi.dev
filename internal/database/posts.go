package database

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Post struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Slug        string     `json:"slug" db:"slug"`
	Content     string     `json:"content" db:"content"`
	Active      int        `json:"active" db:"active"`
	CreatedByID int        `json:"created_by_id" db:"created_by_id"`
	PostedAt    *time.Time `json:"posted_at" db:"posted_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (db *DB) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (title, slug, content, active, created_by_id, posted_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	err := db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Slug,
		post.Content,
		post.Active,
		post.CreatedByID,
		post.PostedAt,
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetByID(ctx context.Context, postID int) (*Post, bool, error) {
	query := `
		SELECT title, slug, content, active, created_by_id, posted_at, created_at, updated_at
		FROM posts
		WHERE id=$1
	`

	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	post := Post{ID: postID}

	err := db.QueryRowContext(ctx, query, postID).Scan(
		&post.Title,
		&post.Slug,
		&post.Content,
		&post.Active,
		&post.CreatedByID,
		&post.PostedAt,
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, err
	}

	return &post, true, nil
}
