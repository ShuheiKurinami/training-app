package security

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CSRFトークンを生成する
func GenerateCSRFToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

// CSRFミドルウェア
func CSRF() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストのCSRFトークンを取得
		reqToken := c.GetHeader("X-CSRF-Token")
		cookieToken, err := c.Cookie("csrf_token")

		// 検証
		if c.Request.Method != http.MethodGet && (err != nil || reqToken != cookieToken) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid CSRF token"})
			c.Abort()
			return
		}

		// 次のハンドラーへ
		c.Next()
	}
}

// ログイン時にCSRFトークンを設定
func SetCSRFToken(c *gin.Context) error {
	token, err := GenerateCSRFToken()
	if err != nil {
		return err
	}
	c.SetCookie("csrf_token", token, 3600, "/", "", true, true)
	c.Header("X-CSRF-Token", token)
	return nil
}

// CSRFセットアップ（起動時に確認用）
func SetupCSRF() {
	println("CSRF setup complete")
}
