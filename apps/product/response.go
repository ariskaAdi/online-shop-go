package product

import "time"

type ProductListResponse struct {
	Id    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewProductListResponseFromEntity(products []ProductEntity) []ProductListResponse {
	var productList = []ProductListResponse{}

	for _, p := range products {
		productList = append(productList, ProductListResponse{
			Id:    p.Id,
			SKU:   p.SKU,
			Name:  p.Name,
			Stock: p.Stock,
			Price: p.Price,
		})
	}

	return productList
}

type ProductDetailResponse struct {
	Id        int    `json:"id"`
	SKU       string `json:"sku"`
	Name      string `json:"name"`
	Stock     int16  `json:"stock"`
	Price     int    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}