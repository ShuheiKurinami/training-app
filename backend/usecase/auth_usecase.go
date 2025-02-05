package usecase

import (
	"errors"

	"github.com/ShuheiKurinami/training-app/backend/domain/models"
	"github.com/ShuheiKurinami/training-app/backend/domain/repositories"
)

type AuthUsecase struct {
	AuthRepo repositories.AuthRepository
}

func NewAuthUsecase(authRepo repositories.AuthRepository) *AuthUsecase {
	return &AuthUsecase{AuthRepo: authRepo}
}

func (uc *AuthUsecase) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := uc.AuthRepo.Authenticate(username, password)
	if err != nil {
		return nil, errors.New("authentication failed")
	}
	return user, nil
}
