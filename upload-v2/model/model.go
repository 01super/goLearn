package model

import (
	"errors"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
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

// InfoAdd 向数据库插入数据
func InfoAdd(mod *Info) error {
	result, err := Db.Exec("insert into info (`name`, path, unix, note) values (?,?,?,?)", mod.Name, mod.Path, mod.Unix, mod.Note)
	if err != nil {
		return err
	}
	id, _ := result.LastInsertId()
	if id < 1 {
		return errors.New("插入失败")
	}
	return nil
}

// InfoGet 获取对应id的信息
func InfoGet(id int64) (Info, error) {
	mod := Info{}
	err := Db.Get(&mod, "select * from info where id = ?", id)
	return mod, err
}

// InfoList 获取列表
func InfoList() ([]Info, error) {
	mod := make([]Info, 0, 8)
	err := Db.Select(&mod, "select * from info")
	return mod, err
}
