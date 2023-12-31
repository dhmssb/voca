package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

func (q *Queries) AddAccountBalance(ctx context.Context, arg presentations.AddAccountBalanceParams) (entity.Account, error) {

	query := `
	UPDATE accounts
	SET balance = balance + $1
	WHERE id = $2
	RETURNING id, owner, balance, currency, created_at
	`

	row := q.db.QueryRowContext(ctx, query, arg.Amount, arg.ID)
	var i entity.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) CreateAccount(ctx context.Context, arg presentations.CreateAccountParams) (entity.Account, error) {

	query := `
	INSERT INTO accounts (
	  owner,
	  balance,
	  currency
	) VALUES (
	  $1, $2, $3
	) RETURNING id, owner, balance, currency, created_at`

	row := q.db.QueryRowContext(ctx, query, arg.Owner, arg.Balance, arg.Currency)
	var i entity.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) GetAccount(ctx context.Context, id int64) (entity.Account, error) {

	query := `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
	`
	row := q.db.QueryRowContext(ctx, query, id)
	var i entity.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (entity.Account, error) {

	query := `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE id = $1 LIMIT 1
	FOR NO KEY UPDATE`

	row := q.db.QueryRowContext(ctx, query, id)
	var i entity.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) ListAccounts(ctx context.Context, arg presentations.ListAccountsParams) ([]entity.Account, error) {

	query := `
	SELECT id, owner, balance, currency, created_at FROM accounts
	WHERE owner = $1
	ORDER BY id
	LIMIT $2
	OFFSET $3
	`

	rows, err := q.db.QueryContext(ctx, query, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.Account{}
	for rows.Next() {
		var i entity.Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
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

func (q *Queries) UpdateAccount(ctx context.Context, arg presentations.UpdateAccountParams) (entity.Account, error) {

	query := `
	UPDATE accounts
	SET balance = $2
	WHERE id = $1
	RETURNING id, owner, balance, currency, created_at
	`
	row := q.db.QueryRowContext(ctx, query, arg.ID, arg.Balance)
	var i entity.Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {

	query := `
	DELETE FROM accounts
	WHERE id = $1
	`
	_, err := q.db.ExecContext(ctx, query, id)
	return err
}
