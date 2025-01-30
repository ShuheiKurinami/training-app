// backend/domain/repositories/user_repository.go

package repositories

import "github.com/ShuheiKurinami/training-app/backend/domain/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]models.User, error)
}
