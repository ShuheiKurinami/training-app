// backend/config/config.go

package config

import (
	"log"
	"os"
)

func LoadEnv() {
	// 環境変数が未設定の場合はデフォルト値を設定
	if os.Getenv("DB_HOST_NODE1") == "" {
		os.Setenv("DB_HOST_NODE1", "localhost")
	}
	if os.Getenv("DB_PORT_NODE1") == "" {
		os.Setenv("DB_PORT_NODE1", "5432")
	}
	if os.Getenv("DB_NAME_NODE1") == "" {
		os.Setenv("DB_NAME_NODE1", "training_app_db_node1")
	}

	if os.Getenv("DB_HOST_NODE2") == "" {
		os.Setenv("DB_HOST_NODE2", "localhost")
	}
	if os.Getenv("DB_PORT_NODE2") == "" {
		os.Setenv("DB_PORT_NODE2", "5433")
	}
	if os.Getenv("DB_NAME_NODE2") == "" {
		os.Setenv("DB_NAME_NODE2", "training_app_db_node2")
	}

	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8080")
	}

	log.Println("Environment variables loaded")
}
