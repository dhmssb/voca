package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

func (q *Queries) CreateEntry(ctx context.Context, arg presentations.CreateEntryParams) (entity.Entry, error) {

	query := `
	INSERT INTO entries (
	  account_id,
	  amount
	) VALUES (
	  $1, $2
	) RETURNING id, account_id, amount, created_at
	`
	row := q.db.QueryRowContext(ctx, query, arg.AccountID, arg.Amount)
	var i entity.Entry
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) ListEntries(ctx context.Context, arg presentations.ListEntriesParams) ([]entity.Entry, error) {
	query := `
	SELECT id, account_id, amount, created_at FROM entries
	WHERE account_id = $1
	ORDER BY id
	LIMIT $2
	OFFSET $3
	`

	rows, err := q.db.QueryContext(ctx, query, arg.AccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.Entry{}
	for rows.Next() {
		var i entity.Entry
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
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
