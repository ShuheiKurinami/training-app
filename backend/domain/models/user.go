// backend/domain/models/user.go

package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // パスワードはJSONに含めない（編集時のみ）
}
