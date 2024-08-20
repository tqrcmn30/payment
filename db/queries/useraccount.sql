-- name: CreateUserAccount :one
INSERT INTO user_accounts (usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING usac_id, usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id;

-- name: GetUserAccountByID :one
SELECT usac_id, usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id
FROM user_accounts
WHERE usac_id = $1;

-- name: ListUserAccounts :many
SELECT usac_id, usac_account_no, usac_balance, usac_created_on, usac_buty_id, usac_user_id
FROM user_accounts
ORDER BY usac_created_on DESC;

-- name: DeleteUserAccount :exec
DELETE FROM user_accounts WHERE usac_id = $1;

-- name: UpdateUserAccount :exec
UPDATE user_accounts
SET usac_account_no = $2, usac_balance = $3, usac_created_on = $4, usac_buty_id = $5, usac_user_id = $6
WHERE usac_id = $1;