-- +goose Up

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT,
    google_id TEXT UNIQUE
);

-- +goose Down

DROP TABLE users;