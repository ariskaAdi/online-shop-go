package product

import (
	"ariskaAdi-online-shop/infra/response"
	"context"
)

type Repository interface{
	CreateProduct(ctx context.Context, model ProductEntity) (err error)
	GetAllProductPagination(ctx context.Context, model ProductPaginationEntity) (products []ProductEntity, err error)
	GetProductBySKU(ctx context.Context, sku string) (product ProductEntity, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{repo: repo}
}

func (s service) CreateProduct(ctx context.Context, req CreateProductRequestPayload) (err error) {
	productEntity := newProductFromCreateProductRequest(req)
	
	if err = productEntity.Validate(); err != nil {
		return
	}

	if err = s.repo.CreateProduct(ctx, productEntity); err != nil {
		return
	}

	return
}

func (s service) ListProduct(ctx context.Context, req ListProductRequestPayload) (product []ProductEntity, err error) {
	pagination := NewProductPaginationFromListProductRequest(req)

	product, err = s.repo.GetAllProductPagination(ctx, pagination)
	if err != nil {
		if err == response.ErrNotFound {
			return []ProductEntity{}, nil
		}

		return
	}

	if len(product) == 0 {
		return []ProductEntity{}, nil
	}
	return
}

func (s service) GetPoductDetail(ctx context.Context, sku string) (model ProductEntity, err error) {
	model, err = s.repo.GetProductBySKU(ctx, sku)
	if err != nil {
		return
	}
	return
}