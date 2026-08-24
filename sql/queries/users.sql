-- name: CreateUser :one
INSERT INTO users (id, username, hashed_password, email, created_at, updated_at)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1;

-- name: UpdateUserById :one
UPDATE users
SET username = $1, hashed_password = $2, email = $3, updated_at = NOW()
WHERE id = $4
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
