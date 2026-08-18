-- name: CreateWeekday :one
INSERT INTO weekday (id, name, day_order, created_at, updated_at, calendar_id, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW(),
    $3,
    $4
)
RETURNING *;

-- name: DeleteWeekday :exec
DELETE FROM weekday
WHERE id = $1;

-- name: GetWeekdaysByCalendarId :many
SELECT *
FROM weekday
WHERE calendar_id = $1
ORDER BY day_order ASC;