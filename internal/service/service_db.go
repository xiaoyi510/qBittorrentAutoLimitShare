package service

import (
	"QbittorrentAutoLimitShare/internal/lib"
	"log"
)

var ServiceDb = &serviceDb{}

type serviceDb struct {
	Db *lib.DbSqlite
}

func (this *serviceDb) CheckUpdate() {
	// 创建表
	exec, err := this.Db.GetDb().Exec("CREATE TABLE if not exists \"tracker\" (\n  \"id\" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,\n  \"tracker\" TEXT,\n  \"limit_rate\" integer,\n  \"limit_time\" integer,\n  \"status\" integer\n)")
	if err != nil {
		log.Panicln(err)
	}
	log.Println(exec.RowsAffected())

}

func (this *serviceDb) Init() {
	ServiceDb.Db = &lib.DbSqlite{}
	ServiceDb.Db.Init().GetDb()
	ServiceDb.CheckUpdate()
}
