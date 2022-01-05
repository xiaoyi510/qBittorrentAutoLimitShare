package handler

import (
	"QbittorrentAutoLimitShare/internal/model/qbit/torrents"
	"QbittorrentAutoLimitShare/internal/service"
	"github.com/spf13/viper"
	"log"
	"math"
	"strings"
	"time"
)

var HandleCron = &handleCron{}

type handleCron struct {
	conf *viper.Viper
}

func (this *handleCron) initConf() {
	v := viper.New()
	v.AddConfigPath("./conf")
	v.SetConfigType("yaml")
	//v.SetConfigFile("app")
	v.SetConfigName("app")

	//missing configuration for 'configPath'
	//////以下方式2
	//v.AddConfigPath("./conf")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("读取配置失败: ", err.Error())
	}
	v.WatchConfig()
	this.conf = v
}
func (this *handleCron) Run() {
	this.initConf()
	// 设置配置项
	service.ServiceCron.SetConf(this.conf)

	log.Println("初始化完成,开始检测Cookie")

	// 判断缓存是否可用
	if service.ServiceCron.CheckCookie() == false {
		log.Println("Cookie无效 开始使用账号密码登录")
		err := service.ServiceCron.Login()
		if err != nil {
			log.Panicln("登录失败", err)
		}
	}
	log.Println("登录完成")

	// 登录成功
	// 开始监听
	go func() {
		for {
			log.Println("开始扫描")
			// 获取信任的tracker列表
			trustTrackers := this.conf.Get("trust_trackers").(string)
			// 获取种子监控最长时间
			skillMaxCompleteTime := 24 * 60 * 60 * this.conf.GetInt("qbit_skill_max_complete_time")
			// 获取配置的上传限制
			SeedingTimeLimit := this.conf.GetInt("qbit_upload_time")
			RatioLimit := this.conf.GetFloat64("qbit_upload_radio")

			trustTrackersArr := strings.Split(trustTrackers, " ")

			// 获取种子tracker列表
			err, s := service.ServiceCron.GetSync().Maindata()
			if err != nil {
				return
			}

			// 定义需要处理的种子列表
			var hashes []string

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
					//var hashNames string
					for _, v2 := range v {
						// 查看种子是否超过监控时间
						a := int(time.Now().Unix() - int64(skillMaxCompleteTime))
						if s.Torrents[v2].CompletionOn > a {
							hashes = append(hashes, v2)
							//hashNames = hashNames + "\r\n" + s.Torrents[v2].Name
						}
					}
				}

			}

			//>> 通知设置分享率
			if len(hashes) > 0 {
				// 去重
				hashes = RemoveRepeatedElement(hashes)

				// 分段处理
				list := ArraySplit(hashes, 6)

				for _, v := range list {
					err, _ := service.ServiceCron.GetTorrents().SetShareLimits(v, torrents.ApiTorrentSetShareLimitsReq{
						SeedingTimeLimit: SeedingTimeLimit,
						RatioLimit:       RatioLimit,
					})
					if err != nil {
						log.Println("设置分享率失败", err.Error())
					} else {
						log.Println("设置分享率成功 共:", len(v))
					}
				}
			}

			// 间隔扫描时间
			limitTime := this.conf.GetDuration("qbit_scan_time")
			if limitTime == 0 {
				this.conf.Set("qbit_scan_time", "10")
				err := this.conf.WriteConfig()
				if err != nil {
					log.Println("写入配置失败,请注意配置文件权限")
				}
				limitTime = 10
			}
			log.Println("开始等待下一轮 等待", int(limitTime), "s")
			time.Sleep(time.Second * limitTime)
		}
	}()
	for {
		// 外部监听避免掉线
		time.Sleep(time.Second * 10)
	}
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

func ArraySplit(arr []string, splitCount int) [][]string {
	arrLen := int(math.Ceil(float64(len(arr)) / float64(splitCount)))
	i := 1
	var ret [][]string
	for i <= arrLen {
		// 计算开始裁剪位置
		left := (i - 1) * splitCount
		right := left + splitCount
		// 如果右边限定值超出总数 则修改到最右侧
		if right > len(arr) {
			right = len(arr)
		}
		ret = append(ret, arr[left:right])
		i++
	}
	return ret
}
