package http

type OptionsSetReq struct {
	ScanTime      int              `json:"scan_time,omitempty"`
	CheckTimeType int              `json:"check_time_type,omitempty"`
	Server        OptionsSetServer `json:"server"`
}

type OptionsSetServer struct {
	ServerUrl      string `json:"url,omitempty"`
	ServerPort     string `json:"port,omitempty"`
	ServerSSL      string `json:"ssl,omitempty"`
	ServerUsername string `json:"username,omitempty"`
	ServerPassword string `json:"password,omitempty"`
}
