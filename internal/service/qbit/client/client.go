package client

import (
	"QbittorrentAutoLimitShare/internal/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	url2 "net/url"
	"strings"
)

type QbitClient struct {
	conf       model.ModelQbit
	httpClient *http.Client
	jar        http.CookieJar
	cookies    []*http.Cookie
}

// Init 初始化
func (this *QbitClient) Init(serverUrl string, serverPort string, isSSL bool) *QbitClient {
	this.conf.ServerUrl = serverUrl
	if serverPort == "" {
		if isSSL {
			this.conf.ServerPort = "443"
		} else {
			this.conf.ServerPort = "80"
		}
	} else {
		this.conf.ServerPort = serverPort
	}
	this.conf.IsSSL = isSSL

	this.jar, _ = cookiejar.New(&cookiejar.Options{})
	this.httpClient = &http.Client{
		Jar: this.jar,
	}
	return this
}

func (this *QbitClient) GetScheme() string {
	if this.conf.IsSSL {
		return "https"
	}
	return "http"
}

func (this *QbitClient) GetHost() string {
	httpScheme := this.GetScheme()
	sprintf := fmt.Sprintf("%s://%s:%s/api/v2",
		httpScheme,
		this.conf.ServerUrl,
		this.conf.ServerPort,
	)
	return sprintf
}

func (this *QbitClient) GetHostHeader() string {
	sprintf := fmt.Sprintf("%s:%s",
		this.conf.ServerUrl,
		this.conf.ServerPort,
	)
	return sprintf
}

func (this *QbitClient) Get(url string, data interface{}) (string, int) {
	dataMap, _ := json.Marshal(data)
	fmt.Println(string(dataMap))

	url = this.GetHost() + "/" + url
	//fmt.Println("[GET] Url:" + url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err.Error(), 0
	}

	response, err := this.httpClient.Do(request)
	if err != nil {
		if response == nil {
			return err.Error(), -1
		}
		return err.Error(), response.StatusCode
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error(), response.StatusCode
	}

	return string(body), response.StatusCode
}

func (this *QbitClient) Post(url string, data interface{}) (string, int) {
	tmpDataMap, _ := json.Marshal(data)

	dataMap := make(map[string]interface{})
	err := json.Unmarshal(tmpDataMap, &dataMap)
	if err != nil {
		return err.Error(), 0
	}

	url = this.GetHost() + "/" + url
	//fmt.Println("[POST] Url:" + url)

	///////////////////////////

	form := url2.Values{}
	for k, v := range dataMap {
		form.Add(k, fmt.Sprintf("%v", v))
	}
	//fmt.Println(form.Encode())
	///////////////////////////
	request, err := http.NewRequest("POST", url, strings.NewReader(form.Encode()))
	if err != nil {
		return err.Error(), 0
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := this.httpClient.Do(request)
	if err != nil {
		if response == nil {
			return err.Error(), -1
		}
		return err.Error(), response.StatusCode
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error(), response.StatusCode
	}
	this.cookies = response.Cookies()
	return string(body), response.StatusCode
}

func (this *QbitClient) GenHashs(hashs []string) string {
	return strings.Join(hashs, "|")
}

// SetCookie 设置Cookie
func (this *QbitClient) SetCookie(cookie []*http.Cookie) *QbitClient {
	cookieUrl := &url2.URL{
		Scheme:      this.GetScheme(),
		Opaque:      "",
		User:        nil,
		Host:        this.GetHostHeader(),
		Path:        "",
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    "",
		Fragment:    "",
		RawFragment: "",
	}
	this.httpClient.Jar.SetCookies(cookieUrl, cookie)
	return this
}

func (this *QbitClient) GetCookie() string {
	var ret string
	for _, v := range this.cookies {
		ret += v.Name + "=" + v.Value + "; "
	}
	return ret
}
