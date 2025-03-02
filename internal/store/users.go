package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/karbasia/karbasi.dev/internal/database"
)

type UserCore struct {
	ID       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
}

type User struct {
	ID             int        `json:"id" db:"id"`
	FullName       string     `json:"full_name" db:"full_name"`
	Email          string     `json:"email" db:"email"`
	HashedPassword string     `json:"-" db:"hashed_password"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitzero" db:"deleted_at"`
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	query := `
		INSERT INTO users (full_name, email, hashed_password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
		`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.FullName,
		user.Email,
		user.HashedPassword,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) Update(ctx context.Context, user *User) error {
	query := `
		UPDATE users SET (full_name, email, hashed_password, deleted_at) =
		($1, $2, $3, $4)
		WHERE id = $5
		RETURNING created_at, updated_at;
	`

	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.FullName,
		user.Email,
		user.HashedPassword,
		user.DeletedAt,
		user.ID,
	).Scan(
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetByID(ctx context.Context, id int) (*User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	user := User{ID: id}

	query := `SELECT full_name, email, hashed_password, created_at, updated_at, deleted_at FROM users WHERE id = $1`
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&user.FullName,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	return &user, true, err
}

func (s *UserStore) GetByEmail(ctx context.Context, email string) (*User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	user := User{Email: email}

	query := `SELECT id, full_name, hashed_password, created_at, updated_at, deleted_at FROM users WHERE email = $1`

	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, false, nil
	}

	return &user, true, err
}

func (s *UserStore) GetAll(ctx context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, database.DefaultTimeout)
	defer cancel()

	query := `SELECT id, full_name, email, created_at, updated_at FROM users WHERE deleted_at IS NULL`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(
			&u.ID,
			&u.FullName,
			&u.Email,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
