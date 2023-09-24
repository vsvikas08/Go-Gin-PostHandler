-- +goose Up

CREATE TABLE users (
    email TEXT PRIMARY KEY,
    id INT NOT NULL
);

-- +goose Down
DROP TABLE users;