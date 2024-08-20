package services

import "github.com/jackc/pgx/v5/pgxpool"

type ServiceManager struct {
	*EntityService
	*UserAccountService
	*PaymentTransactionService
}

func NewServiceManager(dbConn *pgxpool.Conn) *ServiceManager {
	return &ServiceManager{
		EntityService:             NewEntityService(dbConn),
		UserAccountService:        NewUserAccountService(dbConn),
		PaymentTransactionService: NewPaymentTransactionService(dbConn),
	}
}
