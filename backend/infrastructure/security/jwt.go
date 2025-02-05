package security

import (
	"time"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// 秘密鍵（環境変数から取得するのがベスト）
var jwtSecret = []byte("your-secret-key")

// クレーム（トークン内のデータ構造）
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// JWTを生成する
func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 1日間有効
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWTを検証する
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// JWTセットアップ（起動時に確認用）
func SetupJWT() {
	fmt.Println("JWT setup complete")
}
