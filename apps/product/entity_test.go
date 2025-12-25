package product

import (
	"ariskaAdi-online-shop/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)



func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := ProductEntity{
			Name: "honda civic",
			Stock: 10,
			Price: 100_000,
		}

		err := product.Validate()
		require.Nil(t, err)
	})

	t.Run("product required", func(t *testing.T) {
		product := ProductEntity{
			Name: "",
			Stock: 10,
			Price: 100_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

	t.Run("product required", func(t *testing.T) {
		product := ProductEntity{
			Name: "a",
			Stock: 10,
			Price: 100_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})

	t.Run("price required", func(t *testing.T) {
		product := ProductEntity{
			Name: "baju",
			Stock: 10,
			Price: 0,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})

	t.Run("stock required", func(t *testing.T) {
		product := ProductEntity{
			Name: "baju",
			Stock: 0,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})
}