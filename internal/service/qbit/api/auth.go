package api

import (
	"QbittorrentAutoLimitShare/internal/model/qbit/auth"
	"QbittorrentAutoLimitShare/internal/service/qbit/client"
	"fmt"
)

type Auth struct {
	client *client.QbitClient
}

func (this *Auth) SetClient(client *client.QbitClient) *Auth {
	this.client = client
	return this
}

func (this *Auth) Login(username string, password string) (string, string) {
	// 调用API进行登录
	req := auth.ApiAuthLoginReq{
		Username: username,
		Password: password,
	}
	res, _ := this.client.Post("auth/login", req)

	fmt.Println("登录结果", res)
	return res, this.client.GetCookie()
}
