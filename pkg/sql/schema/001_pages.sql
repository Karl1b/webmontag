-- +goose Up
CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    domain TEXT UNIQUE,
    searchDeep int,
    maxSearchDeep int
);

-- +goose Down
DROP TABLE pages;

