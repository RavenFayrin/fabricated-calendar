-- name: CreateUniverse :one
INSERT INTO universe (id, name, description, created_at, updated_at, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW(),
    $3
)
RETURNING *;

-- name: GetUniversesByUserId :many
SELECT * FROM universe
WHERE user_id = $1;