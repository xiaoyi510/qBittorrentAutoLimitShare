package qbit

import (
	"QbittorrentAutoLimitShare/internal/service/qbit/api"
	"QbittorrentAutoLimitShare/internal/service/qbit/client"
)

var Qbit = qbitApi{}

type qbitApi struct {
	Client   client.QbitClient
	Auth     api.Auth
	App      api.App
	Torrents api.Torrents
	Sync     api.Sync
}
