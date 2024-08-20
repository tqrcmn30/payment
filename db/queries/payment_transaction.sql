-- name: CreatePaymentTransaction :one
INSERT INTO payment_transactions (
    patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id;

-- name: GetPaymentTransactionByNo :one
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
FROM payment_transactions
WHERE patrx_no = $1;

-- name: ListPaymentTransactions :many
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
FROM payment_transactions
ORDER BY patrx_created_on DESC;

-- name: DeletePaymentTransaction :exec
DELETE FROM payment_transactions WHERE patrx_no = $1;

-- name: UpdatePaymentTransaction :exec
UPDATE payment_transactions
SET patrx_created_on = $2, patrx_debet = $3, patrx_credit = $4, patrx_notes = $5, patrx_acctno_from = $6, patrx_acctno_to = $7, patrx_patrx_ref = $8, patrx_traty_id = $9
WHERE patrx_no = $1;