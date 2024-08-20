// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: payment_transaction.sql

package db

import (
	"context"

	
)

const createPaymentTransaction = `-- name: CreatePaymentTransaction :one
INSERT INTO payment_transactions (
    patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
`

type CreatePaymentTransactionParams struct {
	PatrxNo         string         `json:"patrx_no"`
	PatrxCreatedOn  *string    `json:"patrx_created_on"`
	PatrxDebet      *int32 `json:"patrx_debet"`
	PatrxCredit     *int32 `json:"patrx_credit"`
	PatrxNotes      *string        `json:"patrx_notes"`
	PatrxAcctnoFrom *string        `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string        `json:"patrx_acctno_to"`
	PatrxPatrxRef   *string        `json:"patrx_patrx_ref"`
	PatrxTratyID    *int32         `json:"patrx_traty_id"`
}

func (q *Queries) CreatePaymentTransaction(ctx context.Context, arg CreatePaymentTransactionParams) (*PaymentTransaction, error) {
	row := q.db.QueryRow(ctx, createPaymentTransaction,
		arg.PatrxNo,
		arg.PatrxCreatedOn,
		arg.PatrxDebet,
		arg.PatrxCredit,
		arg.PatrxNotes,
		arg.PatrxAcctnoFrom,
		arg.PatrxAcctnoTo,
		arg.PatrxPatrxRef,
		arg.PatrxTratyID,
	)
	var i PaymentTransaction
	err := row.Scan(
		&i.PatrxNo,
		&i.PatrxCreatedOn,
		&i.PatrxDebet,
		&i.PatrxCredit,
		&i.PatrxNotes,
		&i.PatrxAcctnoFrom,
		&i.PatrxAcctnoTo,
		&i.PatrxPatrxRef,
		&i.PatrxTratyID,
	)
	return &i, err
}

const deletePaymentTransaction = `-- name: DeletePaymentTransaction :exec
DELETE FROM payment_transactions WHERE patrx_no = $1
`

func (q *Queries) DeletePaymentTransaction(ctx context.Context, patrxNo string) error {
	_, err := q.db.Exec(ctx, deletePaymentTransaction, patrxNo)
	return err
}

const getPaymentTransactionByNo = `-- name: GetPaymentTransactionByNo :one
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
FROM payment_transactions
WHERE patrx_no = $1
`

func (q *Queries) GetPaymentTransactionByNo(ctx context.Context, patrxNo string) (*PaymentTransaction, error) {
	row := q.db.QueryRow(ctx, getPaymentTransactionByNo, patrxNo)
	var i PaymentTransaction
	err := row.Scan(
		&i.PatrxNo,
		&i.PatrxCreatedOn,
		&i.PatrxDebet,
		&i.PatrxCredit,
		&i.PatrxNotes,
		&i.PatrxAcctnoFrom,
		&i.PatrxAcctnoTo,
		&i.PatrxPatrxRef,
		&i.PatrxTratyID,
	)
	return &i, err
}

const listPaymentTransactions = `-- name: ListPaymentTransactions :many
SELECT patrx_no, patrx_created_on, patrx_debet, patrx_credit, patrx_notes, patrx_acctno_from, patrx_acctno_to, patrx_patrx_ref, patrx_traty_id
FROM payment_transactions
ORDER BY patrx_created_on DESC
`

func (q *Queries) ListPaymentTransactions(ctx context.Context) ([]*PaymentTransaction, error) {
	rows, err := q.db.Query(ctx, listPaymentTransactions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*PaymentTransaction
	for rows.Next() {
		var i PaymentTransaction
		if err := rows.Scan(
			&i.PatrxNo,
			&i.PatrxCreatedOn,
			&i.PatrxDebet,
			&i.PatrxCredit,
			&i.PatrxNotes,
			&i.PatrxAcctnoFrom,
			&i.PatrxAcctnoTo,
			&i.PatrxPatrxRef,
			&i.PatrxTratyID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePaymentTransaction = `-- name: UpdatePaymentTransaction :exec
UPDATE payment_transactions
SET patrx_created_on = $2, patrx_debet = $3, patrx_credit = $4, patrx_notes = $5, patrx_acctno_from = $6, patrx_acctno_to = $7, patrx_patrx_ref = $8, patrx_traty_id = $9
WHERE patrx_no = $1
`

type UpdatePaymentTransactionParams struct {
	PatrxNo         string         `json:"patrx_no"`
	PatrxCreatedOn  *string    `json:"patrx_created_on"`
	PatrxDebet      *int32 `json:"patrx_debet"`
	PatrxCredit     *int32 `json:"patrx_credit"`
	PatrxNotes      *string        `json:"patrx_notes"`
	PatrxAcctnoFrom *string        `json:"patrx_acctno_from"`
	PatrxAcctnoTo   *string        `json:"patrx_acctno_to"`
	PatrxPatrxRef   *string        `json:"patrx_patrx_ref"`
	PatrxTratyID    *int32         `json:"patrx_traty_id"`
}

func (q *Queries) UpdatePaymentTransaction(ctx context.Context, arg UpdatePaymentTransactionParams) error {
	_, err := q.db.Exec(ctx, updatePaymentTransaction,
		arg.PatrxNo,
		arg.PatrxCreatedOn,
		arg.PatrxDebet,
		arg.PatrxCredit,
		arg.PatrxNotes,
		arg.PatrxAcctnoFrom,
		arg.PatrxAcctnoTo,
		arg.PatrxPatrxRef,
		arg.PatrxTratyID,
	)
	return err
}