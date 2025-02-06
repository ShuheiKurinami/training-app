package controllers

import (
	"net/http"
	"strconv"

	"github.com/ShuheiKurinami/training-app/backend/domain/models"
	"github.com/ShuheiKurinami/training-app/backend/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUC *usecase.UserUsecase
}

func NewUserController(uc *usecase.UserUsecase) *UserController {
	return &UserController{
		UserUC: uc,
	}
}

// ユーザー作成
func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.UserUC.RegisterUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// パスワードを含めない
	user.Password = ""
	ctx.JSON(http.StatusCreated, user)
}

// ユーザー取得
func (c *UserController) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.UserUC.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// パスワードを含めない
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

// ユーザー更新
func (c *UserController) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	user.ID = id

	if err := c.UserUC.UpdateUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	// パスワードを含めない
	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

// ユーザー削除
func (c *UserController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := c.UserUC.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

// 全ユーザー取得
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.UserUC.FetchAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// パスワードを含めない
	for i := range users {
		users[i].Password = ""
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdatePassword(ctx *gin.Context) {
	// URLからユーザーID取得
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// JSONからパスワードを取得
	var input struct {
		NewPassword string `json:"new_password"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// パスワード変更処理
	err = uc.UserUC.ChangePassword(id, input.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	ctx.Status(http.StatusOK)
}