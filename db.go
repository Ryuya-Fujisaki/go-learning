package main

import (
	"database/sql"
	"log"

	- "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
func InitDB() {
	var err error
	dsn := "username:password@tcp(127.0.0.1:3306)/yourdbname?parseTime=true"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("DB接続失敗:", err)
	} 

	if err = DB.Ping(); err != nil {
		log.Fatal("DBに接続できません:", err)
	}
}