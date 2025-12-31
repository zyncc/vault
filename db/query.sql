-- name: GetAllPasswords :many
SELECT * FROM password_store;

-- name: InsertIntoPasswordStore :exec
INSERT INTO password_store (id, domain, email, password)
VALUES (?, ?, ?, ?);

-- name: FindPasswordUsingDomain :one
SELECT * FROM password_store WHERE domain = ?;

-- name: CreateMasterPassword :exec
INSERT INTO master_password (password, salt)
VALUES (?, ?);

-- name: GetMasterPassword :one
SELECT * FROM master_password;