package presentations

import "voca/internal/entity"

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type ListTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    entity.Tansfer `json:"transfer"`
	FromAccount entity.Account `json:"from_account"`
	ToAccount   entity.Account `json:"to_account"`
	FromEntry   entity.Entry   `json:"from_entry"`
	ToEntry     entity.Entry   `json:"to_entry"`
}
