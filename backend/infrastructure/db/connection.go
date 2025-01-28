// backend/infrastructure/db/connection.go

package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	DB *sql.DB
}

var (
	DBNode1 *DBConnection
	DBNode2 *DBConnection
)

func Init() {
	DBNode1 = connectToDB("NODE1")
	DBNode2 = connectToDB("NODE2")
}

func connectToDB(node string) *DBConnection {
	host := os.Getenv(fmt.Sprintf("DB_HOST_%s", node))
	port := os.Getenv(fmt.Sprintf("DB_PORT_%s", node))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv(fmt.Sprintf("DB_NAME_%s", node))

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Database connection failed for %s: %v", node, err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Database ping failed for %s: %v", node, err)
	}

	log.Printf("Connected to %s database successfully", node)

	return &DBConnection{DB: db}
}
