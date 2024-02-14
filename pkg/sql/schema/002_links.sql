-- +goose Up
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    domain TEXT UNIQUE    
);

-- +goose Down
DROP TABLE links;