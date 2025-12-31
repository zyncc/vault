-- name: GetAllPasswords :many
SELECT * FROM password_store;

-- name: InsertIntoPasswordStore :exec
INSERT INTO password_store (id, domain, password)
VALUES (?, ?, ?);

-- name: FindPasswordUsingDomain :one
SELECT * FROM password_store WHERE domain = ?;