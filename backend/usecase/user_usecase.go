// backend/usecase/user_usecase.go

package usecase

import (
	"github.com/ShuheiKurinami/training-app/backend/domain/models"
	"github.com/ShuheiKurinami/training-app/backend/domain/repositories"
	"github.com/ShuheiKurinami/training-app/backend/infrastructure/security"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepo: repo,
	}
}

func (u *UserUsecase) RegisterUser(user *models.User) error {
	// パスワードのハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return u.UserRepo.CreateUser(user)
}

func (u *UserUsecase) GetUser(id int) (*models.User, error) {
	return u.UserRepo.GetUserByID(id)
}

func (u *UserUsecase) UpdateUser(user *models.User) error {
	if user.Password != "" {
		// パスワードが更新されている場合はハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return u.UserRepo.UpdateUser(user)
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.UserRepo.DeleteUser(id)
}

func (u *UserUsecase) FetchAllUsers() ([]models.User, error) {
	return u.UserRepo.GetAllUsers()
}

func (u *UserUsecase) ChangePassword(id int, newPassword string) error {
	hashed, err := security.HashPassword(newPassword)
	if err != nil {
		return err
	}
	return u.UserRepo.UpdateUserPassword(id, hashed)
}