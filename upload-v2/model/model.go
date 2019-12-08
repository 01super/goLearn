package model

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Db 数据库操作句柄
var Db *sqlx.DB

func init() {
	db1, err := sqlx.Open(`mysql`, `root:88888888@tcp(127.0.0.1:3306)/imageWiew?charset=utf8&parseTime=true`)
	if err != nil {
		log.Fatalln(err)
	}
	err = db1.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	Db = db1
}

// Info 数据库结构
type Info struct {
	ID, Unix         int64
	Name, Path, Note string
}
