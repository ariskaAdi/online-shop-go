package auth

import (
	"ariskaAdi-online-shop/external/database"
	"ariskaAdi-online-shop/infra/response"
	"ariskaAdi-online-shop/internal/config"
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
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

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Name: "adi", 
		Email: fmt.Sprintf("%v@gmail.com", uuid.NewString()), 
		Password: "rahasia",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err) 
}

func TestRegister_Failed(t *testing.T) {
	t.Run("email already exist", func(t *testing.T) {
		// prepare for duplicate email
		email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
		req := RegisterRequestPayload{
			Name: "adi",
			Email: email,
			Password: "rahasia",
		} 
		err := svc.register(context.Background(), req)
		require.Nil(t, err) 

		// end prepare
		err = svc.register(context.Background(), req)
		require.NotNil(t, err) 
		require.Equal(t, response.ErrEmailAlreadyExist, err) 
	})
	
}

func TestLoginSuccess(t *testing.T) {
	email := fmt.Sprintf("%v@gmail.com", uuid.NewString())
	pass := "rahasia"
	req := RegisterRequestPayload{
		Name: "adi", 
		Email:  email,
		Password: pass,
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email: email,
		Password: pass,
	}
	token, err := svc.login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)
	log.Println(token)
}