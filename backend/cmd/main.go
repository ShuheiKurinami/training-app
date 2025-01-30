package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/ShuheiKurinami/training-app/backend/config"
	"github.com/ShuheiKurinami/training-app/backend/infrastructure/db"
	"github.com/ShuheiKurinami/training-app/backend/infrastructure/router"
	"github.com/ShuheiKurinami/training-app/backend/interface/controllers"
	"github.com/ShuheiKurinami/training-app/backend/interface/repositories"
	"github.com/ShuheiKurinami/training-app/backend/usecase"
)

func main() {
	// .env の読み込み
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	// 環境変数の読み込み
	config.LoadEnv()

	// データベース接続の初期化
	db.Init() // ← エラーハンドリングを削除
	log.Println("Database connected successfully")

	// リポジトリの作成
	userRepo := repositories.NewPostgresUserRepository()

	// ユースケースの作成
	userUC := usecase.NewUserUsecase(userRepo)

	// コントローラの作成
	userController := controllers.NewUserController(userUC)

	// ルーターの設定
	r := router.SetupRoutes(userController)

	// サーバー起動
	port := getPort()
	log.Printf("Server is running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
