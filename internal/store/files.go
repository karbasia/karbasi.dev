package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/karbasia/karbasi.dev/internal/database"
)

type File struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Content   []byte     `json:"content,omitzero" db:"content"`
	CreatedAt *time.Time `json:"created_at,omitzero" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitzero" db:"updated_at"`
	DeletedAt *string    `json:"deleted_at,omitzero" db:"deleted_at"`
}

type FileStore struct {
	db *sql.DB
}

func (s *FileStore) Create(ctx context.Context, file *File) error {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()
	query := `
		INSERT INTO files (name, content)
		VALUES($1, $2)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		file.Name,
		file.Content,
	).Scan(
		&file.ID,
		&file.CreatedAt,
		&file.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *FileStore) GetAll(ctx context.Context) ([]File, error) {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()
	query := `
		SELECT id, name, created_at, updated_at, deleted_at
		FROM files
	`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	files := []File{}
	for rows.Next() {
		var f File
		err := rows.Scan(
			&f.ID,
			&f.Name,
			&f.CreatedAt,
			&f.UpdatedAt,
			&f.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}

func (s *FileStore) GetByID(ctx context.Context, id int) (*File, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	file := File{ID: id}

	query := `
		SELECT name, content, created_at, updated_at, deleted_at
		FROM files
		WHERE id = $1
	`
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&file.Name,
		&file.Content,
		&file.CreatedAt,
		&file.UpdatedAt,
		&file.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	return &file, true, err
}

func (s *FileStore) GetByName(ctx context.Context, name string) (*File, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	file := File{Name: name}

	query := `
		SELECT id, content, created_at, updated_at, deleted_at
		FROM files
		WHERE name = $1
	`
	err := s.db.QueryRowContext(ctx, query, name).Scan(
		&file.ID,
		&file.Content,
		&file.CreatedAt,
		&file.UpdatedAt,
		&file.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	return &file, true, err
}
