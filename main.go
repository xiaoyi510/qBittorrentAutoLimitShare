package main

import (
	"QbittorrentAutoLimitShare/internal/cmd"
	"QbittorrentAutoLimitShare/internal/consts"
	"QbittorrentAutoLimitShare/internal/service"
)

func main() {
	service.ServiceDb.Init()
	cmd.PublicWeb.Run()
	cmd.PublicWeb.Run()
	//cmd.HandleCron.Run()
	<-consts.Sig

}
