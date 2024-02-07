package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDb() *sql.DB {
	dsn := "root:root@tcp(localhost:3306)/db_crud"
	db,err := sql.Open("mysql",dsn)
	if err != nil {
		panic(err)
	}

	return db
}