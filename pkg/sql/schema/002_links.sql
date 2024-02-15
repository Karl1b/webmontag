-- +goose Up
CREATE TABLE links (
    id SERIAL PRIMARY KEY,
    statusCode INT,
    domain TEXT UNIQUE    
);

-- +goose Down
DROP TABLE links;