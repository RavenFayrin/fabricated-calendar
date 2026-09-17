-- name: CreateEra :one
INSERT INTO era (id, name, shorthand, start_year, description, created_at, updated_at, calendar_id, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    NOW(),
    NOW(),
    $5,
    $6
)
RETURNING *;

-- name: GetErasByCalendarId :many
SELECT *
FROM era
WHERE calendar_id = $1
ORDER BY start_year ASC;

-- name: UpdateEraById :one
UPDATE era
SET name = $1, shorthand = $2, start_year = $3, description = $4, updated_at = NOW()
WHERE id = $5
RETURNING *;

-- name: DeleteMonth :exec
DELETE FROM era
WHERE id = $1;