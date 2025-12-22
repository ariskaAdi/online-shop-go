package auth

import (
	"ariskaAdi-online-shop/infra/response"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthEntity(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "adi@gmail.com",
			Password: "rahasia",
		}

		err := authEntity.Validate()
		require.Nil(t, err)
	})

		t.Run("email is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "",
			Password: "rahasia",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailRequired, err)
	})

		t.Run("email is invalid", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "adi.com",
			Password: "rahasia",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailInvalid, err)
	})

		t.Run("password is required", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "adi@gmail.com",
			Password: "",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordRequired, err)
	})

		t.Run("password is must be at least 6 characters", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "adi@gmail.com",
			Password: "ra",
		}

		err := authEntity.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPasswordInvalid, err)
	})
}

func TestEncryptPassword(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		authEntity := AuthEntity{
			Email: "adi@gmail.com",
			Password: "rahasia",
		}

		err := authEntity.EncryptPassword(bcrypt.DefaultCost)
		require.Nil(t, err)

		log.Printf("%+v\n", authEntity)
		
	})
}