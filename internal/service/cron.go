package service

import (
	"QbittorrentAutoLimitShare/internal/service/qbit"
	"QbittorrentAutoLimitShare/internal/service/qbit/api"
	"QbittorrentAutoLimitShare/internal/service/qbit/client"
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

var ServiceCron *serviceCron

func init() {
	ServiceCron = &serviceCron{}
}

type serviceCron struct {
	client *client.QbitClient
	conf   *viper.Viper
}

func (this *serviceCron) Init() {
	this.client = qbit.Qbit.Client.Init(this.conf.GetString("qbit_url"), this.conf.GetString("qbit_port"), this.conf.GetString("qbit_ssl") == "1")
}

// 第一次运行 同步所有的数据
func (this *serviceCron) Login() error {
	//>> 登录获取Cookie
	res := this.GetAuth().Login(this.conf.GetString("qbit_username"), this.conf.GetString("qbit_password"))
	if res == "Ok." {
		return nil
	}
	return errors.New("登录失败:" + res)
}

func (this *serviceCron) GetApp() *api.App {
	return (&qbit.Qbit.App).SetClient(this.client)
}
func (this *serviceCron) GetTorrents() *api.Torrents {
	return (&qbit.Qbit.Torrents).SetClient(this.client)
}

func (this *serviceCron) GetAuth() *api.Auth {
	return (&qbit.Qbit.Auth).SetClient(this.client)
}
func (this *serviceCron) GetSync() *api.Sync {
	return (&qbit.Qbit.Sync).SetClient(this.client)
}

func (this *serviceCron) SetConf(conf *viper.Viper) *serviceCron {
	this.conf = conf
	this.Init()
	return this
}

// CheckCookie 判断cookie是否可用
func (this *serviceCron) CheckCookie() bool {
	cookie := this.conf.Get("qbit_cookie").(string)
	if len(cookie) <= 0 {
		return false
	}
	this.SetCookie(cookie)

	err, _ := this.GetApp().Version()
	if err != nil {
		return false
	}
	return true
}

// 解析Cookie
func (this *serviceCron) parseCookie(cookie string) []*http.Cookie {
	rawCookies := cookie
	rawRequest := fmt.Sprintf("GET / HTTP/1.0\r\nCookie: %s\r\nHost: "+this.client.GetHostHeader()+"\r\n\r\n", rawCookies)

	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(rawRequest)))

	if err == nil {
		return req.Cookies()
	}
	return nil
}
func (this *serviceCron) SetCookie(cookie string) {
	this.client.SetCookie(this.parseCookie(cookie))
}
