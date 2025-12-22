package auth

import (
	"ariskaAdi-online-shop/external/database"
	"ariskaAdi-online-shop/internal/config"
	"context"
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

func TestRegister_Success(t *testing.T) {
	req := RegisterRequestPayload{
		Name: "adi", 
		Email: "adi@gmail.com", 
		Password: "rahasia",
	}
	err := svc.register(context.Background(), req)
	require.Nil(t, err) 
}

// 52