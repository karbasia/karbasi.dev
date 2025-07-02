-- +goose Up
CREATE TABLE posts_to_tags (
    post_id INTEGER NOT NULL REFERENCES posts(id),
    tag_id INTEGER NOT NULL REFERENCES tags(id),
    PRIMARY KEY(post_id, tag_id)
);

-- +goose Down
DROP TABLE posts_to_tags;