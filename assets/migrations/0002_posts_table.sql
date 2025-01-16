-- +goose Up
CREATE TABLE posts (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    content TEXT NOT NULL,
    active INTEGER NOT NULL DEFAULT 0,
    created_by_id INTEGER NOT NULL,
    posted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN KEY(created_by_id) REFERENCES users(id)
);
-- +goose StatementBegin
CREATE TRIGGER update_posts_updated_at
AFTER UPDATE ON posts
WHEN OLD.updated_at <> CURRENT_TIMESTAMP
BEGIN
    UPDATE posts
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TABLE posts;