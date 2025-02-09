package repositories

import (
	"database/sql"
	"errors"
	"github.com/ShuheiKurinami/training-app/backend/domain/models"
	"github.com/ShuheiKurinami/training-app/backend/domain/repositories"
	"github.com/ShuheiKurinami/training-app/backend/infrastructure/db"
	"golang.org/x/crypto/bcrypt"
)

type PostgresAuthRepository struct{}

func NewPostgresAuthRepository() repositories.AuthRepository {
	return &PostgresAuthRepository{}
}

// 認証処理（データベースからユーザー情報を取得）
func (r *PostgresAuthRepository) Authenticate(username, password string) (*models.User, error) {
	var user models.User
	var hashedPassword string

	// まずは DBNode1 で検索
	err := db.DBNode1.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
		Scan(&user.ID, &user.Username, &hashedPassword)

	// ユーザーが見つからない場合は DBNode2 で再検索
	if err == sql.ErrNoRows {
		err = db.DBNode2.DB.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).
			Scan(&user.ID, &user.Username, &hashedPassword)
	}

	// どちらのノードにもユーザーが存在しない
	if err == sql.ErrNoRows {
		return nil, errors.New("invalid credentials")
	} else if err != nil {
		return nil, err
	}

	// パスワードの検証
	if !comparePasswords(hashedPassword, password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// パスワードの比較関数（bcryptを使用）
func comparePasswords(hashedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
