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
	DeletedAt *string    `json:"deleted_at,omitzero" db:"deleted_at"`
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

func (s *TagStore) GetAll(ctx context.Context, showDeleted bool, params PaginationParams) (PaginatedResult[Tag], error) {
	filterParam := ""
	if !showDeleted {
		filterParam = "WHERE deleted_at IS NULL"
	}

	countQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM tags
		%s
	`, filterParam)

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	var totalItems int
	err := s.db.QueryRowContext(ctx, countQuery).Scan(&totalItems)
	if err != nil {
		return PaginatedResult[Tag]{}, err
	}

	offset := (params.Page - 1) * params.PageSize
	totalPages := 0
	if params.PageSize > 0 {
		totalPages = (totalItems + params.PageSize - 1) / params.PageSize
	}

	query := fmt.Sprintf(`
		SELECT id, name, created_at, updated_at, deleted_at
		FROM tags
		%s
		ORDER BY name
		LIMIT $1 OFFSET $2
	`, filterParam)

	rows, err := s.db.QueryContext(ctx, query, params.PageSize, offset)
	if err != nil {
		return PaginatedResult[Tag]{}, err
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
			return PaginatedResult[Tag]{}, err
		}
		tags = append(tags, t)
	}

	return PaginatedResult[Tag]{
		Items: tags,
		Pagination: PaginationMeta{
			Page:       params.Page,
			PageSize:   params.PageSize,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}, nil
}

func (s *TagStore) GetAllByPostCount(ctx context.Context, params PaginationParams) (PaginatedResult[Tag], error) {
	countQuery := `
		SELECT COUNT(*)
		FROM tags t
		INNER JOIN posts_to_tags pt ON t.id = pt.tag_id
		INNER JOIN posts p ON p.id = pt.post_id
		WHERE p.deleted_at IS NULL
			AND t.deleted_at IS NULL
		GROUP BY t.id
	`
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	countRows, err := s.db.QueryContext(ctx, countQuery)
	if err != nil {
		return PaginatedResult[Tag]{}, err
	}
	defer countRows.Close()

	totalItems := 0
	for countRows.Next() {
		totalItems++
	}

	offset := (params.Page - 1) * params.PageSize
	totalPages := 0
	if params.PageSize > 0 {
		totalPages = (totalItems + params.PageSize - 1) / params.PageSize
	}

	query := `
		SELECT t.id, t.name, t.created_at, t.updated_at, t.deleted_at, COUNT(p.id) AS post_count
		FROM tags t
		LEFT JOIN posts_to_tags pt ON t.id = pt.tag_id
		LEFT JOIN posts p ON p.id = pt.post_id
		WHERE p.deleted_at IS NULL
			AND t.deleted_at IS NULL
		GROUP BY t.id
		ORDER BY COUNT(p.id) DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := s.db.QueryContext(ctx, query, params.PageSize, offset)
	if err != nil {
		return PaginatedResult[Tag]{}, err
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
			return PaginatedResult[Tag]{}, err
		}
		tags = append(tags, t)
	}

	return PaginatedResult[Tag]{
		Items: tags,
		Pagination: PaginationMeta{
			Page:       params.Page,
			PageSize:   params.PageSize,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}, nil
}
