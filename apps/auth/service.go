package auth

import (
	"ariskaAdi-online-shop/infra/response"
	"ariskaAdi-online-shop/internal/config"
	"context"
)

type Repository interface{
	GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFormRegisterRequest(req)

	if err = authEntity.Validate(); err != nil {
		return 
	}

	if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}

	if model.IsExist() {
		return response.ErrEmailAlreadyExist
	}

	return s.repo.CreateAuth(ctx, authEntity)
}