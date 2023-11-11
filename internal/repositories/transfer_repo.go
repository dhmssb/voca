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

func (q *Queries) GetTransfer(ctx context.Context, id int64) (entity.Tansfer, error) {

	query := `
	SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
	WHERE id = $1 LIMIT 1
	`

	row := q.db.QueryRowContext(ctx, query, id)
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

func (q *Queries) ListTransfers(ctx context.Context, arg presentations.ListTransfersParams) ([]entity.Tansfer, error) {

	query := `
	SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
	WHERE 
		from_account_id = $1 OR
		to_account_id = $2
	ORDER BY id
	LIMIT $3
	OFFSET $4
	`
	rows, err := q.db.QueryContext(ctx, query,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.Tansfer{}
	for rows.Next() {
		var i entity.Tansfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
