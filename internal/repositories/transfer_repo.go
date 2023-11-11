package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

func (q *Queries) CreateTransfer(ctx context.Context, arg presentations.CreateTransferParams) (entity.Tansfer, error) {
	query := `
	INSERT INTO transfers (
	  from_account_id,
	  to_account_id,
	  amount
	) VALUES (
	  $1, $2, $3
	) RETURNING id, from_account_id, to_account_id, amount, created_at
	`

	row := q.db.QueryRowContext(ctx, query, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i entity.Tansfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}
