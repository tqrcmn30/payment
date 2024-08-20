package services

import (
	db "payment/db/sqlc"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserAccountService struct {
	*db.Queries
}

// constructor
func NewUserAccountService(dbConn *pgxpool.Conn) *UserAccountService {
	return &UserAccountService{
		Queries: db.New(dbConn),
	}
}
