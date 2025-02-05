package repositories

import "github.com/ShuheiKurinami/training-app/backend/domain/models"

type AuthRepository interface {
	Authenticate(username, password string) (*models.User, error)
}
