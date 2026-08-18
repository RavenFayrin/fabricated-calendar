-- name: CreateCalendar :one
INSERT INTO calendar (id, name, description, created_at, updated_at, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW(),
    $3
)
RETURNING *;

-- name: DeleteCalendar :exec
DELETE FROM calendar
WHERE id = $1;

-- name: GetCalendarsByUserId :many
SELECT * FROM calendar
WHERE user_id = $1;