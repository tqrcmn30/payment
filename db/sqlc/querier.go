// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreateEntity(ctx context.Context, entyName string) (*EntityType, error)
	CreatePaymentTransaction(ctx context.Context, arg CreatePaymentTransactionParams) (*PaymentTransaction, error)
	CreateUserAccount(ctx context.Context, arg CreateUserAccountParams) (*UserAccount, error)
	DeleteEntity(ctx context.Context, entyID int32) error
	DeletePaymentTransaction(ctx context.Context, patrxNo string) error
	DeleteUserAccount(ctx context.Context, usacID int32) error
	FindAllEntity(ctx context.Context) ([]*EntityType, error)
	FindEntityById(ctx context.Context, entyID int32) (*EntityType, error)
	GetPaymentTransactionByNo(ctx context.Context, patrxNo string) (*PaymentTransaction, error)
	GetUserAccountByID(ctx context.Context, usacID int32) (*UserAccount, error)
	ListPaymentTransactions(ctx context.Context) ([]*PaymentTransaction, error)
	ListUserAccounts(ctx context.Context) ([]*UserAccount, error)
	UpdateEntity(ctx context.Context, arg UpdateEntityParams) (*EntityType, error)
	UpdatePaymentTransaction(ctx context.Context, arg UpdatePaymentTransactionParams) error
	UpdateUserAccount(ctx context.Context, arg UpdateUserAccountParams) error
}

var _ Querier = (*Queries)(nil)
