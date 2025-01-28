// backend/interface/repositories/user_repository_impl.go

package repositories

import (
	"math"
	"strconv"
	"training-app/backend/domain/models"
	"training-app/backend/infrastructure/db"

	"github.com/pkg/errors"
)

type PostgresUserRepository struct{}

func NewPostgresUserRepository() UserRepository {
	return &PostgresUserRepository{}
}

func (r *PostgresUserRepository) getDBConnection(userID int) *db.DBConnection {
	// ユーザーIDに基づいてデータベースノードを選択
	// 例: 1-100はノード1, 101-200はノード2
	if userID <= 100 {
		return db.DBNode1
	} else if userID <= 200 {
		return db.DBNode2
	}
	// さらにノードを追加する場合はここに条件を追加
	// デフォルトはノード1を返す
	return db.DBNode1
}

func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	// 新規作成時はuserIDはまだ決まっていないため、ノードを選択する基準を変更
	// ここでは単純にノード1を使用
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
	err := db.DBNode1.DB.QueryRow(query, user.Username, user.Email, user.Password).Scan(&user.ID)
	return err
}

func (r *PostgresUserRepository) GetUserByID(id int) (*models.User, error) {
	dbConn := r.getDBConnection(id)
	user := &models.User{}
	query := `SELECT id, username, email, password FROM users WHERE id = $1`
	err := dbConn.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return user, nil
}

func (r *PostgresUserRepository) UpdateUser(user *models.User) error {
	dbConn := r.getDBConnection(user.ID)
	query := `UPDATE users SET username=$1, email=$2, password=$3 WHERE id=$4`
	res, err := dbConn.DB.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}
	if rowsAffected == 0 {
		return errors.New("no user found to update")
	}
	return nil
}

func (r *PostgresUserRepository) DeleteUser(id int) error {
	dbConn := r.getDBConnection(id)
	query := `DELETE FROM users WHERE id=$1`
	res, err := dbConn.DB.Exec(query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}
	if rowsAffected == 0 {
		return errors.New("no user found to delete")
	}
	return nil
}

func (r *PostgresUserRepository) GetAllUsers() ([]models.User, error) {
	// 全ユーザーを取得する場合、全ノードからデータを取得し統合
	var allUsers []models.User

	// ノード1から取得
	usersNode1, err := r.getUsersFromDB(db.DBNode1.DB)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get users from node1")
	}
	allUsers = append(allUsers, usersNode1...)

	// ノード2から取得
	usersNode2, err := r.getUsersFromDB(db.DBNode2.DB)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get users from node2")
	}
	allUsers = append(allUsers, usersNode2...)

	// さらにノードがある場合は同様に取得
	return allUsers, nil
}

func (r *PostgresUserRepository) getUsersFromDB(dbConn *db.DBConnection.DB) ([]models.User, error) {
	query := `SELECT id, username, email FROM users`
	rows, err := dbConn.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query users")
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan user")
		}
		users = append(users, user)
	}
	return users, nil
}
