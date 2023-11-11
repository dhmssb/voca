package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg presentations.AddAccountBalanceParams) (entity.Account, error)
	CreateAccount(ctx context.Context, arg presentations.CreateAccountParams) (entity.Account, error)
	CreateEntry(ctx context.Context, arg presentations.CreateEntryParams) (entity.Entry, error)
	CreateProduct(ctx context.Context, arg presentations.CreateProductParams) (entity.Product, error)
	CreateTransfer(ctx context.Context, arg presentations.CreateTransferParams) (entity.Tansfer, error)
	CreateUser(ctx context.Context, arg presentations.CreateUserParams) (entity.User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteProduct(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (entity.Account, error)
	GetAccountForUpdate(ctx context.Context, id int64) (entity.Account, error)
	GetEntry(ctx context.Context, id int64) (entity.Entry, error)
	GetProduct(ctx context.Context, id int64) (entity.Product, error)
	GetProductForUpdate(ctx context.Context, id int64) (entity.Product, error)
	GetTransfer(ctx context.Context, id int64) (entity.Tansfer, error)
	GetUser(ctx context.Context, username string) (entity.User, error)
	ListAccounts(ctx context.Context, arg presentations.ListAccountsParams) ([]entity.Account, error)
	ListEntries(ctx context.Context, arg presentations.ListEntriesParams) ([]entity.Entry, error)
	ListProducts(ctx context.Context, arg presentations.ListProductsParams) ([]entity.Product, error)
	ListTransfers(ctx context.Context, arg presentations.ListTransfersParams) ([]entity.Tansfer, error)
	UpdateAccount(ctx context.Context, arg presentations.UpdateAccountParams) (entity.Account, error)
	UpdateProductQuantity(ctx context.Context, arg presentations.UpdateProductQuantityParams) (entity.Product, error)
}

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg presentations.TransferTxParams) (presentations.TransferTxResult, error)
}

var _ Querier = (*Queries)(nil)
