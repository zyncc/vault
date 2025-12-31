-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS master_password (
    password TEXT PRIMARY KEY,
    salt TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS master_password;
-- +goose StatementEnd
