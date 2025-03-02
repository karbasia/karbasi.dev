package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
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
	PostedAt    *time.Time `json:"posted_at,omitzero" db:"posted_at"`
	CreatedBy   UserCore   `json:"created_by"`
	Tags        []Tag      `json:"tags"`
	CreatedAt   *time.Time `json:"created_at,omitzero" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitzero" db:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitzero" db:"deleted_at"`
}

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *Post) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		INSERT INTO posts (title, slug, content, active, created_by_id, posted_at)
		VALUES ($1, $2, $3, $4, $5, $6) 
		RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	err = tx.QueryRowContext(
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

	err = associateTags(ctx, tx, post.ID, post.Tags)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) Update(ctx context.Context, post *Post) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `
		UPDATE posts SET (title, slug, content, active, posted_at, deleted_at) =
		($1, $2, $3, $4, $5, $6)
		WHERE id = $7
		RETURNING created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	err = tx.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Slug,
		post.Content,
		post.Active,
		post.PostedAt,
		post.DeletedAt,
		post.ID,
	).Scan(
		&post.CreatedAt,
		&post.UpdatedAt,
	)
	if err != nil {
		return err
	}

	err = removeTags(ctx, tx, post.ID, post.Tags)
	if err != nil {
		return err
	}

	err = associateTags(ctx, tx, post.ID, post.Tags)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *PostStore) GetBySlug(ctx context.Context, slug string) (*Post, bool, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.content, p.active, p.created_by_id, p.posted_at, p.created_at, p.updated_at, p.deleted_at, 
			u.id, u.full_name,
			json_group_array(
				json_object('id', t.id, 'name', t.name)
			) filter (
				where t.id IS NOT NULL
			) AS tags
		FROM posts p
		INNER JOIN users u ON p.created_by_id = u.id
		LEFT JOIN posts_to_tags pt ON p.id = pt.post_id
		LEFT JOIN tags t ON pt.tag_id = t.id
		WHERE p.slug=$1
		GROUP BY p.id
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	post := Post{}
	var tagData string

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
		&post.DeletedAt,
		&post.CreatedBy.ID,
		&post.CreatedBy.FullName,
		&tagData,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	} else if err != nil {
		return nil, false, err
	}
	err = json.Unmarshal([]byte(tagData), &post.Tags)
	if err != nil {
		return nil, false, err
	}

	return &post, true, nil
}

func (s *PostStore) GetAllByTag(ctx context.Context, tagName string) ([]Post, error) {
	query := `
		SELECT p.id, p.title, p.slug, p.active, p.created_by_id, p.posted_at, p.created_at, p.updated_at, u.id, u.full_name,
			json_group_array(
				json_object('id', t.id, 'name', t.name)
			) filter (
				where t.id IS NOT NULL
			) AS tags
		FROM posts p
		INNER JOIN users u ON p.created_by_id = u.id
		INNER JOIN posts_to_tags pt ON p.id = pt.post_id
		INNER JOIN tags t ON pt.tag_id = t.id
		LEFT JOIN posts_to_tags pt2 ON p.id = pt2.post_id
		LEFT JOIN tags t2 ON pt2.tag_id = t2.id
		WHERE t2.name=$1 AND p.deleted_at IS NULL
		GROUP BY p.id
		ORDER BY p.posted_at DESC
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, tagName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []Post
	for rows.Next() {
		p := Post{}
		var tagData string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&p.Active,
			&p.CreatedByID,
			&p.PostedAt,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.CreatedBy.ID,
			&p.CreatedBy.FullName,
			&tagData,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(tagData), &p.Tags)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (s *PostStore) GetAll(ctx context.Context, showDeleted bool) ([]Post, error) {
	filterParam := ""
	if !showDeleted {
		filterParam = "WHERE p.deleted_at IS NULL"
	}
	query := fmt.Sprintf(`
		SELECT p.id, p.title, p.slug, p.active, p.created_by_id, p.posted_at, p.created_at, p.updated_at, u.id, u.full_name, p.deleted_at,
			json_group_array(
				json_object('id', t.id, 'name', t.name)
			) filter (
				where t.id IS NOT NULL
			) AS tags
		FROM posts p
		INNER JOIN users u ON p.created_by_id = u.id
		LEFT JOIN posts_to_tags pt ON p.id = pt.post_id
		LEFT JOIN tags t ON pt.tag_id = t.id
		%s
		GROUP BY p.id
		ORDER BY p.posted_at DESC
	`, filterParam)

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
		var tagData string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&p.Active,
			&p.CreatedByID,
			&p.PostedAt,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.CreatedBy.ID,
			&p.CreatedBy.FullName,
			&p.DeletedAt,
			&tagData,
		)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(tagData), &p.Tags)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func removeTags(ctx context.Context, tx *sql.Tx, postID int, tags []Tag) error {
	query := `
		DELETE FROM posts_to_tags
		WHERE post_id = $1
			AND tag_id NOT IN (
				SELECT json_extract(json_each.value, '$.id')
				FROM json_each($2)
			);
	`

	data, err := json.Marshal(tags)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	_, err = tx.ExecContext(
		ctx,
		query,
		postID,
		string(data),
	)
	if err != nil {
		return err
	}

	return nil
}

func associateTags(ctx context.Context, tx *sql.Tx, postID int, tags []Tag) error {
	query := `
		INSERT INTO posts_to_tags(post_id, tag_id)
			SELECT $1, id
			FROM tags
			WHERE id IN (
				SELECT json_extract(json_each.value, '$.id')
				FROM json_each($2)
			)
		ON CONFLICT(post_id, tag_id) DO NOTHING;
	`

	data, err := json.Marshal(tags)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	_, err = tx.ExecContext(
		ctx,
		query,
		postID,
		string(data),
	)
	if err != nil {
		return err
	}

	return nil
}
