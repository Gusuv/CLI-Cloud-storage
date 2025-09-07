-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username TEXT NOT NULL UNIQUE,
                       email TEXT NOT NULL UNIQUE,
                       password_hash TEXT NOT NULL,
                       created TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE files (
                       id SERIAL PRIMARY KEY,
                       user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       name TEXT NOT NULL,
                       path TEXT NOT NULL,
                       size BIGINT NOT NULL,
                       mime_type VARCHAR(100),
                       created TIMESTAMPTZ DEFAULT now()
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE files;
DROP TABLE users;