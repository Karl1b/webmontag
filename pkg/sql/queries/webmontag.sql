
-- name: CreatePage :one
INSERT INTO pages (domain,searchDeep,maxSearchDeep) VALUES ($1, $2 , $3) RETURNING *;

/* -- name: CreateLink :many
INSERT INTO links (domain) VALUES ($1) RETURNING id; */

-- name: InsertOrGetLink :one
INSERT INTO links (domain,statusCode)
VALUES ($1,$2)
ON CONFLICT (domain) DO UPDATE 
SET domain = EXCLUDED.domain
RETURNING id;

-- name: DeleteAllPages :exec
DELETE FROM pages;

-- name: GetPageByDomain :many
SELECT id, domain, searchdeep, maxsearchdeep FROM pages WHERE domain LIKE $1;

-- name: GetAllPages :many
SELECT * FROM pages;

-- name: AssociateLinkWithPage :exec
INSERT INTO pageslinks (pages_id, links_id) VALUES ($1, $2);

-- name: GetLinksByPageID :many
SELECT l.id, l.domain, l.statusCode
FROM links l
INNER JOIN pageslinks pl ON l.id = pl.links_id
WHERE pl.pages_id = $1;

-- name: UpsertPage :one
INSERT INTO pages (domain, searchDeep, maxSearchDeep)
VALUES ($1, $2, $3)
ON CONFLICT (domain) DO UPDATE 
SET searchDeep = EXCLUDED.searchDeep, maxSearchDeep = EXCLUDED.maxSearchDeep
RETURNING id;

-- name: InsertPageLinkAssociation :exec
INSERT INTO pageslinks (pages_id, links_id)
VALUES ($1, $2)
ON CONFLICT (pages_id, links_id) DO NOTHING;