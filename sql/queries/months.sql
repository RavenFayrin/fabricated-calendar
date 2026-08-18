-- name: CreateMonth :one
INSERT INTO month (id, name, month_order, days_in_month, created_at, updated_at, calendar_id, user_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    NOW(),
    NOW(),
    $4,
    $5
)
RETURNING *;

-- name: DeleteMonth :exec
DELETE FROM month
WHERE id = $1;

-- name: GetMonthsByCalendarId :many
SELECT *
FROM month
WHERE calendar_id = $1
ORDER BY month_order ASC;
