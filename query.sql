-- name: GetUsers :many
SELECT id,name,email 
FROM users;

-- name: GetUser :one
SELECT id,name,email 
FROM users 
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users(name,email,password)
VALUES($1,$2,$3)
RETURNING id;

-- name: UpdateUser :one
UPDATE users 
SET name = $2, email = $3 
WHERE id = $1 
RETURNING id,name,email;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE id = $1;



-- name: CreateOpinion :one
INSERT INTO opinions (
    user_id, subject_type, subject_id, opinion, rating
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetOpinion :one
SELECT * FROM opinions
WHERE id = $1;

-- name: UpdateOpinion :one
UPDATE opinions
SET opinion = $2,
    rating = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteOpinion :exec
DELETE FROM opinions
WHERE id = $1;




-- name: ListOpinions :many
SELECT *
FROM opinions
WHERE
    -- filtering
    ($1::text IS NULL OR subject_type = $1) AND
    ($2::bigint IS NULL OR subject_id = $2) AND
    ($3::bigint IS NULL OR user_id = $3) AND
    -- searching
    ($4::text IS NULL OR opinion ILIKE '%' || $4 || '%')
ORDER BY created_at DESC
LIMIT $5 OFFSET $6;
