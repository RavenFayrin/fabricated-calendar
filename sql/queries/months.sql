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

-- name: GetMonthsByCalendarId :many
SELECT *
FROM month
WHERE calendar_id = $1
ORDER BY month_order ASC;

-- name: UpdateCalendarById :one
UPDATE month
SET name = $1, month_order = $2, days_in_month = $3, updated_at = NOW()
WHERE id = $4
RETURNING *;

-- name: DeleteMonth :exec
DELETE FROM month
WHERE id = $1;
