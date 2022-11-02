package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// DB is a global variable to hold db connection
var DB *sql.DB

func InitDb(host string, port int, user string, dbname string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	DB = db
	fmt.Println("Database connection successful")
	return nil
}
