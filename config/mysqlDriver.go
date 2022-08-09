package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectionDB() *sql.DB {
	dbUsername := os.Getenv("DB_Username")
	dbPassword := os.Getenv("DB_Password")
	dbPort := os.Getenv("DB_Port")
	dbHost := os.Getenv("DB_Host")
	dbName := os.Getenv("DB_Name")
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName)
	db, err := sql.Open("mysql", dbConnection)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("success")
	}
	return db
}
