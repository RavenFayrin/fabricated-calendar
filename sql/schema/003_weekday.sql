-- +goose Up
CREATE TABLE weekday (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    day_order INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    calendar_id UUID NOT NULL,
    user_id UUID NOT NULL,
    CONSTRAINT fk_calendar_id
        FOREIGN KEY (calendar_id)
        REFERENCES calendar(id)
        ON DELETE CASCADE,
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE weekday;
