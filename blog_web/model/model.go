package model

import (
	"errors"
	"fmt"
	// mysql
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Db 数据库操作句柄
var Db *sqlx.DB

func init() {
	db1, err := sqlx.Open(`mysql`, `root:88888888@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=true`)
	if err != nil {
		log.Fatalln(err)
	}
	err = db1.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	Db = db1
}

// Artical 数据库结构
type Artical struct {
	ID, Time                    int64
	Title, Describe, Main, Type string
}

// ArticalAdd 向数据库插入数据
func ArticalAdd(mod *Artical) error {
	fmt.Println(mod.Title, mod.Describe, mod.Main, mod.Type, mod.Time)
	result, err := Db.Exec("insert into essay (`title`, `describe`, `main`, `type`, `time`) values (?,?,?,?,?)", mod.Title, mod.Describe, mod.Main, mod.Type, mod.Time)
	if err != nil {
		fmt.Println(err)
		return err
	}
	id, _ := result.LastInsertId()
	if id < 1 {
		return errors.New("插入失败")
	}
	return nil
}

// ArticalGet 获取对应id的信息
func ArticalGet(id int64) (Artical, error) {
	mod := Artical{}
	err := Db.Get(&mod, "select * from essay where id = ?", id)
	return mod, err
}

// ArticalList 获取列表
func ArticalList() ([]Artical, error) {
	mod := make([]Artical, 0, 8)
	err := Db.Select(&mod, "select * from essay")
	return mod, err
}

// ArticalDrop delete
func ArticalDrop(id int64) error {
	result, err := Db.Exec("delete from essay where id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows != 1 {
		return errors.New("删除失败")
	}
	return nil
}
