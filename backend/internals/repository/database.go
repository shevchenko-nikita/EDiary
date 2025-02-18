package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/diary")

	if err != nil {
		return nil, fmt.Errorf("Error connecting to DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging DB: %v", err)
	}

	return db, nil
}
