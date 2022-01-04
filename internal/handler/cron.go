package handler

import (
	"QbittorrentAutoLimitShare/internal/service"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

var HandleCron = &handleCron{}

type handleCron struct {
	conf *viper.Viper
}

func (this *handleCron) initConf() {
	v := viper.New()
	v.SetConfigFile("./conf/app.yaml")
	//////以下方式2
	//v.AddConfigPath("./conf")
	//v.SetConfigName("app")
	//v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}
	v.WatchConfig()
	this.conf = v
}
func (this *handleCron) Run() {
	this.initConf()
	// 设置配置项
	service.ServiceCron.SetConf(this.conf)

	// 判断缓存是否可用
	if service.ServiceCron.CheckCookie() == false {
		err := service.ServiceCron.Login()
		if err != nil {
			panic(err)
		}
	}

	// 登录成功
	// 开始监听
	go func() {
		for {
			// 获取信任的tracker列表
			trust_trackers := this.conf.Get("trust_trackers").(string)
			trustTrackersArr := strings.Split(trust_trackers, " ")

			// 获取种子tracker列表

			err, s := service.ServiceCron.GetSync().Maindata()
			if err != nil {
				return
			}

			// 判断哪些tracker不是信任的
			for k, v := range s.Trackers {
				hasTrust := false
				for _, trustUrl := range trustTrackersArr {
					if strings.Index(k, trustUrl) != -1 && len(trustUrl) > 0 {
						hasTrust = true
						break
					}
				}
				if !hasTrust {
					//>> 没有在信任列表
					//>> 判断 v 中是否已处理分享率
					var hashes []string
					for _, v2 := range v {
						hashes = append(hashes, v2)
					}
					//>> 通知设置分享率

					println(k + "\r\n")
				}

			}

			// 间隔扫描时间
			time.Sleep(time.Second * 10)
		}
	}()
	for {
		// 外部监听避免掉线
		time.Sleep(time.Second * 10)
	}
}
