package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ShuheiKurinami/training-app/backend/usecase"
	"github.com/ShuheiKurinami/training-app/backend/infrastructure/security"
)

type AuthController struct {
	AuthUC *usecase.AuthUsecase
}

func NewAuthController(authUC *usecase.AuthUsecase) *AuthController {
	return &AuthController{AuthUC: authUC}
}

// CSRFトークン取得
func (h *AuthController) GetCSRFToken(c *gin.Context) {
	err := security.SetCSRFToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CSRFトークンの生成に失敗しました"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"csrfToken": c.Writer.Header().Get("X-CSRF-Token")})
}

// ログイン処理
func (h *AuthController) Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "リクエストデータが無効です"})
		return
	}

	user, err := h.AuthUC.AuthenticateUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗しました"})
		return
	}

	token, err := security.GenerateJWT(uint(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "JWTトークンの生成に失敗しました"})
		return
	}

	err = security.SetCSRFToken(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "CSRFトークンの設定に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "ログイン成功",
		"token":     token,
		"csrfToken": c.Writer.Header().Get("X-CSRF-Token"),
	})
}

// ログアウト処理
func (h *AuthController) Logout(c *gin.Context) {
	c.SetCookie("csrf_token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{"message": "ログアウト成功"})
}
