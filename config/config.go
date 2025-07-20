package config

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB(){
	var err error
	connStr := "user=postgres password=postgres dbname=go_invoice sslmode=disable"
	DB, err= sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping error: ", err)
	}
}