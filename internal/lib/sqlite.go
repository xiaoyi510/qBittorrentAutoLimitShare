package lib

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DbSqlite struct {
	db *sql.DB
}

func (this *DbSqlite) Init() *DbSqlite {
	var err error
	this.db, err = sql.Open("sqlite3", ".conf/qbit-auto.db")
	if err != nil {
		log.Panicln("数据库打开失败:" + err.Error())
	}
	return this
}

func (this *DbSqlite) GetDb() *sql.DB {
	if this.db == nil {
		this.Init()
	}
	return this.db
}
