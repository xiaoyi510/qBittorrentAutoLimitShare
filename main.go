package main

import (
	"QbittorrentAutoLimitShare/internal/cmd"
	"QbittorrentAutoLimitShare/internal/consts"
)

func main() {
	//service.ServiceDb.Init()
	cmd.PublicWeb.Run()
	//cmd.HandleCron.Run()
	<-consts.Sig

}
