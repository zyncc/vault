-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS password_store (
    id TEXT PRIMARY KEY,
    domain TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS password_store;
-- +goose StatementEnd
