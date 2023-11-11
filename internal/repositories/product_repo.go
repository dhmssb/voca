package repositories

import (
	"context"
	"voca/internal/entity"
	"voca/internal/presentations"
)

func (q *Queries) CreateProduct(ctx context.Context, arg presentations.CreateProductParams) (entity.Product, error) {

	query := `
	INSERT INTO products (
	  user_product,
	  product_name,
	  product_description,
	  quantity,
	  price
	) VALUES (
	  $1, $2, $3, $4, $5
	) RETURNING id, user_product, product_name, product_description, quantity, price
	`

	row := q.db.QueryRowContext(ctx, query,
		arg.UserProduct,
		arg.ProductName,
		arg.ProductDescription,
		arg.Quantity,
		arg.Price,
	)
	var i entity.Product
	err := row.Scan(
		&i.ID,
		&i.UserProduct,
		&i.ProductName,
		&i.ProductDescription,
		&i.Quantity,
		&i.Price,
	)
	return i, err
}

func (q *Queries) DeleteProduct(ctx context.Context, id int64) error {

	query := `
	DELETE FROM products
	WHERE id = $1
	`
	_, err := q.db.ExecContext(ctx, query, id)
	return err
}

func (q *Queries) GetProduct(ctx context.Context, id int64) (entity.Product, error) {

	query := `
	SELECT id, user_product, product_name, product_description, quantity, price FROM products
	WHERE id = $1 LIMIT 1
	`

	row := q.db.QueryRowContext(ctx, query, id)
	var i entity.Product
	err := row.Scan(
		&i.ID,
		&i.UserProduct,
		&i.ProductName,
		&i.ProductDescription,
		&i.Quantity,
		&i.Price,
	)
	return i, err
}

func (q *Queries) GetProductForUpdate(ctx context.Context, id int64) (entity.Product, error) {

	query := `
	SELECT id, user_product, product_name, product_description, quantity, price FROM products
	WHERE id = $1 LIMIT 1
	FOR NO KEY UPDATE
	`

	row := q.db.QueryRowContext(ctx, query, id)
	var i entity.Product
	err := row.Scan(
		&i.ID,
		&i.UserProduct,
		&i.ProductName,
		&i.ProductDescription,
		&i.Quantity,
		&i.Price,
	)
	return i, err
}

func (q *Queries) ListProducts(ctx context.Context, arg presentations.ListProductsParams) ([]entity.Product, error) {

	query := `
	SELECT id, user_product, product_name, product_description, quantity, price FROM products
	WHERE user_product = $1
	ORDER BY id
	LIMIT $2
	OFFSET $3
	`

	rows, err := q.db.QueryContext(ctx, query, arg.UserProduct, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []entity.Product{}
	for rows.Next() {
		var i entity.Product
		if err := rows.Scan(
			&i.ID,
			&i.UserProduct,
			&i.ProductName,
			&i.ProductDescription,
			&i.Quantity,
			&i.Price,
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

func (q *Queries) UpdateProductQuantity(ctx context.Context, arg presentations.UpdateProductQuantityParams) (entity.Product, error) {
	query := `
	UPDATE products
	SET quantity = $2
	WHERE id = $1
	RETURNING id, user_product, product_name, product_description, quantity, price
	`

	row := q.db.QueryRowContext(ctx, query, arg.ID, arg.Quantity)
	var i entity.Product
	err := row.Scan(
		&i.ID,
		&i.UserProduct,
		&i.ProductName,
		&i.ProductDescription,
		&i.Quantity,
		&i.Price,
	)
	return i, err
}
