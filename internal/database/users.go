package database

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID             int        `json:"id" db:"id"`
	FullName       string     `json:"full_name" db:"full_name"`
	Email          string     `json:"email" db:"email"`
	HashedPassword string     `json:"-" db:"hashed_password"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (db *DB) InsertUser(ctx context.Context, fullName, email, hashedPassword string) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `
		INSERT INTO users (full_name, email, hashed_password)
		VALUES ($1, $2, $3)`

	result, err := db.ExecContext(ctx, query, fullName, email, hashedPassword)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), err
}

func (db *DB) GetUserByID(ctx context.Context, id int) (*User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	user := User{ID: id}

	query := `SELECT full_name, email, hashed_password, created_at, updated_at, deleted_at FROM users WHERE id = $1`
	err := db.QueryRowContext(ctx, query, id).Scan(
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

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	user := User{Email: email}

	query := `SELECT id, full_name, hashed_password, created_at, updated_at, deleted_at FROM users WHERE email = $1`

	err := db.QueryRowContext(ctx, query, email).Scan(
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

func (db *DB) UpdateUserHashedPassword(ctx context.Context, id int, hashedPassword string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `UPDATE users SET hashed_password = $1 WHERE id = $2`

	_, err := db.ExecContext(ctx, query, hashedPassword, id)
	return err
}

func (db *DB) GetAllUsers(ctx context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	query := `SELECT id, full_name, email, created_at, updated_at FROM users WHERE deleted_at IS NULL`
	rows, err := db.QueryContext(ctx, query)
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
