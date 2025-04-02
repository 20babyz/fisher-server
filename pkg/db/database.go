package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(user, password, host, dbname string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL: %v", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping MySQL: %v", err)
	}

	DB = db
	return nil
}
