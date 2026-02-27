package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ybuilds/todo-api/internal/utils"
)

var DB *sql.DB

func init() {
	var err error

	connString := fmt.Sprintf(
		`host=%s user=%s password=%s port=%s dbname=%s sslmode=disable`,
		utils.GetFromEnv("DB_HOST"),
		utils.GetFromEnv("DB_USER"),
		utils.GetFromEnv("DB_PASSWORD"),
		utils.GetFromEnv("DB_PORT"),
		utils.GetFromEnv("DB_NAME"),
	)

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("error opening db connection", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("error pinging db", err)
	}

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
}
