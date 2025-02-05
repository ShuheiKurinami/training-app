package repositories

import (
	"errors"

	"github.com/ShuheiKurinami/training-app/backend/domain/models"
	"github.com/ShuheiKurinami/training-app/backend/domain/repositories"
)

type PostgresAuthRepository struct{}

func NewPostgresAuthRepository() repositories.AuthRepository {
	return &PostgresAuthRepository{}
}

// 仮の認証関数（本番環境ではデータベースと連携する）
func (r *PostgresAuthRepository) Authenticate(username, password string) (*models.User, error) {
	// 仮のユーザー情報（DB接続がない場合）
	if username == "admin" && password == "password" {
		return &models.User{
			ID:       1,
			Username: "admin",
			Password: "", // パスワードは返さない
		}, nil
	}
	return nil, errors.New("invalid credentials")
}
