-- +goose Up
CREATE TABLE era (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    shorthand TEXT,
    start_year INTEGER NOT NULL,
    description TEXT,

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
        ON DELETE CASCADE,

    CONSTRAINT uq_calendar_start_year
        UNIQUE (calendar_id, start_year)
);

-- +goose Down
DROP TABLE era;
