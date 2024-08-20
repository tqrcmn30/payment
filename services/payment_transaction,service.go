package services

import (
	db "payment/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentTransactionService struct {
	*db.Queries
}

// constructor
func NewPaymentTransactionService(dbConn *pgxpool.Conn) *PaymentTransactionService {
	return &PaymentTransactionService{
		Queries: db.New(dbConn),
	}
}
