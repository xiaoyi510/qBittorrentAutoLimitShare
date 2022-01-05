package handler

import (
	"QbittorrentAutoLimitShare/internal/model/qbit/torrents"
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
			// 获取种子监控最长时间
			skillMaxCompleteTime := 24 * 60 * 60 * this.conf.GetInt("qbit_skill_max_complete_time")
			// 获取配置的上传限制
			SeedingTimeLimit := this.conf.GetInt("qbit_upload_time")
			RatioLimit := this.conf.GetFloat64("qbit_upload_radio")

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
					var hashNames string
					for _, v2 := range v {
						// 查看种子是否超过监控时间
						a := int(time.Now().Unix() - int64(skillMaxCompleteTime))
						if s.Torrents[v2].CompletionOn > a {
							hashes = append(hashes, v2)
							hashNames = hashNames + "\r\n" + s.Torrents[v2].Name
						}
					}
					//>> 通知设置分享率
					if len(hashes) > 0 {
						err, _ := service.ServiceCron.GetTorrents().SetShareLimits(hashes, torrents.ApiTorrentSetShareLimitsReq{
							SeedingTimeLimit: SeedingTimeLimit,
							RatioLimit:       RatioLimit,
						})
						if err != nil {
							log.Println("设置分享率失败", err.Error())
						} else {
							log.Println("设置分享率成功", hashNames)
						}
					}
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
