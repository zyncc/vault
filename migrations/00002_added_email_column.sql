-- +goose Up
-- +goose StatementBegin
ALTER TABLE password_store ADD COLUMN email TEXT NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE password_store DROP COLUMN email;
-- +goose StatementEnd
