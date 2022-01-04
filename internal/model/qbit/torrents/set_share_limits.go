package torrents

type ApiTorrentsSetShareLimitsReq struct {
	Hashes           string `json:"hashes"`
	RatioLimit       string `json:"ratioLimit"`
	SeedingTimeLimit string `json:"seedingTimeLimit"`
}

type ApiTorrentsSetShareLimitsRes struct {
}
