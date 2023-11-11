package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg presentations.AddAccountBalanceParams) (entity.Account, error)
	CreateAccount(ctx context.Context, arg presentations.CreateAccountParams) (entity.Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (entity.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (entity.Account, error)
	ListAccounts(ctx context.Context, arg presentations.ListAccountsParams) ([]entity.Account, error)
	UpdateAccount(ctx context.Context, arg presentations.UpdateAccountParams) (entity.Account, error)
}

var _ Querier = (*Queries)(nil)
