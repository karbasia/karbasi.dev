-- +goose Up
CREATE TABLE
    files (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL UNIQUE,
        content BLOB NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        deleted_at TEXT
    );

-- +goose StatementBegin
CREATE TRIGGER update_files_updated_at AFTER
UPDATE ON files WHEN OLD.updated_at <> CURRENT_TIMESTAMP BEGIN
UPDATE files
SET
    updated_at = CURRENT_TIMESTAMP
WHERE
    id = OLD.id;

END;

-- +goose StatementEnd
-- +goose Down
DROP TABLE files;