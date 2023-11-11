package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

func (q *Queries) CreateUser(ctx context.Context, arg presentations.CreateUserParams) (entity.User, error) {

	query := `
	INSERT INTO users (
	  username,
	  hashed_password,
	  full_name,
	  email
	) VALUES (
	  $1, $2, $3, $4
	) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
	`
	row := q.db.QueryRowContext(ctx, query,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i entity.User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

func (q *Queries) GetUser(ctx context.Context, username string) (entity.User, error) {
	query := `
	SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
	WHERE username = $1 LIMIT 1
	`

	row := q.db.QueryRowContext(ctx, query, username)
	var i entity.User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
