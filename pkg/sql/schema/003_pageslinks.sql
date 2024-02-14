-- +goose Up
CREATE TABLE pageslinks (
    id SERIAL PRIMARY KEY,
    pages_id INT REFERENCES pages(id) ON DELETE CASCADE,
    links_id INT REFERENCES links(id) ON DELETE CASCADE,
    UNIQUE(pages_id, links_id)
);

-- +goose Down
DROP TABLE pageslinks;