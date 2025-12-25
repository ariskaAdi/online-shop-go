package product

import (
	"ariskaAdi-online-shop/infra/response"
	"time"

	"github.com/google/uuid"
)

type ProductEntity struct {
	Id        int       `db:"id"`
	SKU 	string    `db:"sku"`
	Name      string    `db:"name"`
	Stock     int16     `db:"stock"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	
}

type ProductPaginationEntity struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

// from request object

func NewProductPaginationFromListProductRequest(req ListProductRequestPayload) ProductPaginationEntity {
	req = req.GenerateDefaultValue()
	return ProductPaginationEntity{
		Cursor: req.Cursor,
		Size: req.Size,
	}	
}

func newProductFromCreateProductRequest(req CreateProductRequestPayload) ProductEntity {
	return ProductEntity{
		SKU: uuid.NewString(),
		Name: req.Name,
		Stock: req.Stock,
		Price: req.Price,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}


// Validate global product
func (p ProductEntity) Validate() (err error) {
	if err = p.ValidateName(); err != nil {
		return	
	}
	if err = p.ValidatePrice(); err != nil {
		return	
	}
	if err = p.ValidateStock(); err != nil {
		return	
	}
	return
}

func (p ProductEntity) ValidateName() (err error) {
	if p.Name == "" {
		return response.ErrProductRequired
	}

	if len(p.Name) < 4 {
		return response.ErrProductInvalid
	}
	return
}

func (p ProductEntity) ValidatePrice() (err error) {
	if p.Price <= 0 {
		return response.ErrPriceInvalid
	}
	return
}

func (p ProductEntity) ValidateStock() (err error) {
	if p.Stock <= 0 {
		return response.ErrStockInvalid
	}
	return
}