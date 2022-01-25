package service

import (
	"QbittorrentAutoLimitShare/internal/lib"
	"QbittorrentAutoLimitShare/internal/model/http"
	"log"
)

var ServiceHttpOptions = options{}

type options struct {
}

func (this *options) Set(http.OptionsSetReq) {
	sqlite := lib.DbSqlite{}
	log.Println(sqlite.GetDb().Ping())
}
