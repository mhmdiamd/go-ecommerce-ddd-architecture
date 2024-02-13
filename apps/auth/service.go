package auth

import (
	"context"

	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/infra/response"
	"github.com/mhmdiamd/go-ecommerce-ddd-architecture/internal/config"
)

type Repository interface {
  GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
  CreateAuth(ctx context.Context, model AuthEntity) (err error)
}

type service struct {
  repo Repository
}

func NewService(repo Repository) service {
  return service{
    repo : repo,
  }
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
 
  authEntity := NewFromRegisterRequest(req)

  if err = authEntity.Validate(); err != nil {
    return
  }

  if err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt)); err != nil {
    return
  }

  model, err := s.repo.GetAuthByEmail(ctx, req.Email) 
  if err != nil {
    if err != response.ErrNotFound {
      return 
    }
  }

  if model.IsExists() {
    return response.ErrEmailAlreadyUsed
  }


  return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
  authEntity := NewFromLoginRequest(req)

  if err = authEntity.EmailValidate(); err != nil {
    return 
  }

  if err = authEntity.PasswordValidate(); err != nil {
    return
  }

  model, err := s.repo.GetAuthByEmail(context.Background(), authEntity.Email)
  if err != nil {
    return 
  }

  if err = authEntity.VerifyPasswordFromPlain(model.Password); err != nil {
    err = response.ErrPasswordNotMatch
    return 
  }
 
  token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)

  return
}


















