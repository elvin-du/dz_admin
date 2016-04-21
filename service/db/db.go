package db

import (
	"database/sql"
	"dz_admin/service/log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db  *sql.DB
	dsn = `root:abc@tcp(localhost:3306)/dz?charset=utf8`
)

func Init() {
	var err error
	db, err = sql.Open("mysql", dsn)
	if nil != err {
		log.Fatalln(err)
	}
}

func DB() *sql.DB {
	return db
}
