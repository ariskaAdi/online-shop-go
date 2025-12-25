package product

import (
	"ariskaAdi-online-shop/infra/response"
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) (repository) {
	return repository{db: db}
}

func (r repository) CreateProduct(ctx context.Context, model ProductEntity) (err error) {
	query := `
		INSERT INTO products(
			sku, name, stock, price, created_at, updated_at
		)
		VALUES (
			:sku, :name, :stock, :price, :created_at, :updated_at
		)
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, model); err != nil {
		return
	}

	return
}

func (r repository) GetAllProductPagination(ctx context.Context, model ProductPaginationEntity) (products []ProductEntity, err error) {
	query := `
		SELECT id, sku, name, stock, price, created_at, updated_at
		FROM products
		WHERE id > $1
		ORDER BY id ASC
		LIMIT $2
	`

	err = r.db.SelectContext(ctx, &products, query, model.Cursor, model.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}
	return	
}