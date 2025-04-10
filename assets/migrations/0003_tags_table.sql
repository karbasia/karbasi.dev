-- +goose Up
CREATE TABLE
    tags (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TEXT
    );

-- +goose StatementBegin
CREATE TRIGGER update_tags_updated_at AFTER
UPDATE ON tags WHEN OLD.updated_at <> CURRENT_TIMESTAMP BEGIN
UPDATE tags
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;

-- +goose StatementEnd
-- +goose Down
DROP TABLE tags;