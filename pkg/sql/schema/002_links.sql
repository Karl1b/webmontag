-- +goose Up
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    isValid BOOlEAN,
    domain TEXT UNIQUE    
);

-- +goose Down
DROP TABLE links;