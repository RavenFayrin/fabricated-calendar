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

-- name: GetCalendarsByUserId :many
SELECT * FROM calendar
WHERE user_id = $1;

-- name: UpdateCalendarByID :one
UPDATE calendar
SET name = $1, description = $2, updated_at = NOW()
WHERE id = $3
RETURNING *;

-- name: DeleteCalendar :exec
DELETE FROM calendar
WHERE id = $1;
