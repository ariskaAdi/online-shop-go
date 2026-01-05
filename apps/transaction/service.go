package transaction

import (
	"ariskaAdi-online-shop/infra/response"
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	TransactionDBRepository
	TransactionRepository
	ProductRepository
}

type TransactionDBRepository interface {
	Begin(ctx context.Context) (tx *sqlx.Tx, err error)
	Rollback(ctx context.Context, tx *sqlx.Tx,) ( err error)
	Commit(ctx context.Context, tx *sqlx.Tx,) ( err error)
}

type TransactionRepository interface{
	CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, trx TransactionEntity) (err error)
	GetTransactionsByUserPublicId(ctx context.Context, userPublicId string) (trxs []TransactionEntity, err error)
}

type ProductRepository interface{
	GetPoductBySku(ctx context.Context, productSKU string) (product Product, err error)
	UpdateProductStockWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error)
}

type service struct {
	repo Repository
}

// newService returns a new service instance with given repository.
func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

// CreateTransaction creates a new transaction with given request and product stock.
// It starts a new transaction, creates a new transaction entity, validates the transaction,
// updates the product stock, and commits the transaction to the database.
// If any error occurs, it will rollback the transaction.
func (s service) CreateTransaction(ctx context.Context, req CreateTransactionRequestPayload) (err error) {
	myProduct, err := s.repo.GetPoductBySku(ctx, req.ProductSKU)	
	if err != nil {
		return
	}

	if !myProduct.IsExist(){
		err = response.ErrNotFound
		return	
	}

	trx := NewTransactionFromCreatedRequest(req)
	trx.FormProduct(myProduct).SetPlatformFee(1_000).SetGrandTotal()

	if err = trx.validate(); err != nil {
		return
	}

	if err = trx.validateStock(uint8(myProduct.Stock)); err != nil {
		return
	}
	
	// start tx database
	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	// defer rollback if any error
	defer s.repo.Rollback(ctx, tx)

	if err = s.repo.CreateTransactionWithTx(ctx, tx, trx); err != nil {
		return
	}

	// update current stock
	if err = s.repo.UpdateProductStockWithTx(ctx, tx, myProduct); err != nil {
		return
	}

	// update into database
	if err = s.repo.Commit(ctx, tx); err != nil {
		return
	}
	
	return
}


// Returns transaction histories of a user by user public id. If no transactions found, returns an empty array and nil error. If error occurs, returns an empty array and the error.
func (s service) TransactionsHistories(ctx context.Context, userPublicId string) (trxs []TransactionEntity, err error) {
	trxs, err = s.repo.GetTransactionsByUserPublicId(ctx, userPublicId)
	if err != nil {
		if err == response.ErrNotFound {
			trxs = []TransactionEntity{}
			return trxs, nil
		}

		return
	}

	if len(trxs) == 0 {
		trxs = []TransactionEntity{}
		return
	}
	return
	
}	