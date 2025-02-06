package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/ShuheiKurinami/training-app/backend/interface/controllers"
	"time"
)

func SetupRoutes(userController *controllers.UserController, authController *controllers.AuthController) *gin.Engine {
	router := gin.Default()

	// CORS 設定
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ユーザー関連のルート
	router.POST("/api/users", userController.CreateUser)
	router.GET("/api/users/:id", userController.GetUser)
	router.PUT("/api/users/:id", userController.UpdateUser)
	router.DELETE("/api/users/:id", userController.DeleteUser)
	router.GET("/api/users", userController.GetAllUsers)
	router.PUT("/users/{id}/password", userController.UpdatePassword)

	// 認証関連のルート
	router.POST("/api/auth/login", authController.Login)
	router.POST("/api/auth/logout", authController.Logout)
	router.GET("/api/auth/csrf-token", authController.GetCSRFToken)

	return router
}
