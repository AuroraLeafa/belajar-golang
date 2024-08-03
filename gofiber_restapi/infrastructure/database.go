package app

import (
	"database/sql"
	"gofiber_restapi/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/fiber_todo_db")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(15 * time.Minute)

	return db
}
