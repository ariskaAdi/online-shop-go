package product

import (
	"ariskaAdi-online-shop/external/database"
	"ariskaAdi-online-shop/internal/config"
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service
func init(){
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name: "baju baru",
		Stock: 10,
		Price: 100_000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestGetAllProduct_success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size: 10,
	}

	products, err := svc.ListProduct(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}

func TestGetProductDetail_success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name: "celana baru",
		Stock: 10,
		Price: 100_000,
	}

	ctx := context.Background()

	err := svc.CreateProduct(ctx, req)
	require.Nil(t, err)

	products, err := svc.ListProduct(ctx, ListProductRequestPayload{
		Cursor: 0,
		Size: 10,
	})

	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.GetPoductDetail(ctx, products[0].SKU)
	require.Nil(t, err)
	require.NotNil(t, product)

	log.Printf("%+v", product)
}