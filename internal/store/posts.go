package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/karbasia/karbasi.dev/internal/database"
)

type Post struct {
	ID          int        `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	Slug        string     `json:"slug" db:"slug"`
	Content     string     `json:"content" db:"content"`
	Active      int        `json:"active" db:"active"`
	CreatedByID int        `json:"-" db:"created_by_id"`
	PostedAt    *time.Time `json:"posted_at" db:"posted_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	CreatedBy   UserCore   `json:"created_by"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (title, slug, content, active, created_by_id, posted_at)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	err := s.db.QueryRowContext(
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

func (s *PostStore) GetByID(ctx context.Context, postID int) (*Post, bool, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.content, p.active, p.created_by_id, p.posted_at, p.created_at, p.updated_at, u.id, u.full_name
		FROM posts p
		INNER JOIN users u ON p.created_by_id = u.id
		WHERE p.id=$1
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	post := Post{ID: postID}

	err := s.db.QueryRowContext(ctx, query, postID).Scan(
		&post.Title,
		&post.Slug,
		&post.Content,
		&post.Active,
		&post.CreatedByID,
		&post.PostedAt,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.CreatedBy.ID,
		&post.CreatedBy.FullName,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, err
	}

	return &post, true, nil
}

func (s *PostStore) GetBySlug(ctx context.Context, slug string) (*Post, bool, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.content, p.active, p.created_by_id, p.posted_at, p.created_at, p.updated_at, u.id, u.full_name
		FROM posts p
		INNER JOIN users u ON p.created_by_id = u.id
		WHERE p.slug=$1
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	post := Post{}

	err := s.db.QueryRowContext(ctx, query, slug).Scan(
		&post.ID,
		&post.Title,
		&post.Slug,
		&post.Content,
		&post.Active,
		&post.CreatedByID,
		&post.PostedAt,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.CreatedBy.ID,
		&post.CreatedBy.FullName,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}

	return &post, true, nil
}

func (s *PostStore) GetAll(ctx context.Context) ([]Post, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.content, p.active, p.created_by_id, p.posted_at, p.created_at, p.updated_at, u.id, u.full_name
		FROM posts p
		INNER JOIN users u ON p.created_by_id = u.id
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post
	for rows.Next() {
		p := Post{}
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&p.Content,
			&p.Active,
			&p.CreatedByID,
			&p.PostedAt,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.CreatedBy.ID,
			&p.CreatedBy.FullName,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
