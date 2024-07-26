package handler

import (
	"QbittorrentAutoLimitShare/internal/consts"
	"QbittorrentAutoLimitShare/internal/model/qbit/torrents"
	"QbittorrentAutoLimitShare/internal/service"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"log"
	"os"
	"strconv"
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
	// 初始化配置项
	this.initConf()
	println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>\n\n欢迎使用qBit 自动限制分享率工具 \r\n\r\nBy:包子 Blog:https://blog.52nyg.com\r\n\r\n>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	// 设置配置项
	service.ServiceCron.SetConf(this.conf)

	// 判断是否登录成功
	if service.ServiceCron.IsLogin == false {
		// 没有登录成功则结束
		os.Exit(0)
	}

	// 开始监听
	go func() {
		var runCount int64
		for {
			runCount++
			log.Println("======================================")
			log.Println("开始扫描 第" + strconv.FormatInt(runCount, 10) + "次")
			// 判断配置项是否正确
			if this.conf.Get("trust_trackers") == nil {
				log.Println("请配置 trust_trackers 20s后重试")
				log.Println("======================================\r\n\r\n")
				time.Sleep(20 * time.Second)
				continue
			}
			// 获取信任的tracker列表
			trustTrackers := this.conf.Get("trust_trackers").(string)
			// 分割信任的Tracker
			trustTrackersArr := []string{}
			if len(trustTrackers) > 0 {
				trustTrackersArr = strings.Split(trustTrackers, " ")
			}

			// 获取种子监控最长时间
			skillMaxCompleteTime := 24 * 60 * 60 * this.conf.GetInt("qbit_skip_max_complete_time")

			// 获取种子检测时间类型
			checkTimeType := this.conf.GetInt("qbit_check_time_type")
			if checkTimeType == 0 {
				checkTimeType = consts.SCAN_TIME_TYPE_AC
			}

			// 获取配置的上传限制
			SeedingTimeLimit := this.conf.GetInt("qbit_upload_time") // 上传时间
			RatioLimit := this.conf.GetFloat64("qbit_upload_radio")  // 分享率
			uploadLimit := this.conf.GetInt("qbit_upload_limit")     // 上传速度

			// 获取种子列表
			err, s := service.ServiceCron.GetSync().Maindata()
			if err != nil {
				log.Println("获取种子列表失败", err)
				// 间隔扫描时间
				this.checkScanTime(0)
				// 检查登录状态是否合法
				this.checkLogin()
				continue
			}

			// 定义需要处理的种子列表
			var hashes []string

			trustTrackerMax := this.conf.Get("tracker_max")
			trustTrackerMaxNum := 0
			if trustTrackerMax != nil {
				trustTrackerMaxNum = cast.ToInt(trustTrackerMax)
			}

			trackerTz := make(map[string][]string)
			if trustTrackerMaxNum > 0 {
				// 种子tracker最大数量超过限制处理
				// 整合数据
				for tracker, bittorrentHashs := range s.Trackers {
					for _, hash := range bittorrentHashs {
						trackerTz[hash] = append(trackerTz[hash], tracker)
					}
				}
				for hash, trackers := range trackerTz {
					if len(trackers) > trustTrackerMaxNum {
						hashes = append(hashes, hash)

					}
				}

			}

			// 判断哪些tracker不是信任的
			for k, v := range s.Trackers {
				// 判断是否是信任的 tracker
				hasTrust := false

				// 在配置中循环判断是否有信任的
				for _, trustUrl := range trustTrackersArr {
					if strings.Index(k, trustUrl) != -1 && len(trustUrl) > 0 {
						hasTrust = true
						break
					}
				}
				if len(trustTrackersArr) > 0 && !hasTrust {
					//>> 没有在信任列表
					//>> 判断 v 中是否已处理分享率
					for _, hash := range v {
						hashes = append(hashes, hash)

					}
				}

			}

			hashesChecked := []string{}
			// 去重
			hashes = service.ServiceHelper.RemoveRepeatedElement(hashes)
			for _, hash := range hashes {
				// 如果种子未限制比例|时间 并且有需要设置比例 则处理  已手动限制比例则不处理
				if (s.Torrents[hash].RatioLimit != RatioLimit && RatioLimit != -1) || (s.Torrents[hash].SeedingTimeLimit != SeedingTimeLimit && SeedingTimeLimit != -1) || (s.Torrents[hash].UpLimit != uploadLimit && uploadLimit != 0) {
					// 获取种子最低监控时间
					minScanTime := int(time.Now().Unix() - int64(skillMaxCompleteTime))

					// 获取判断时间
					tmpTime := service.ServiceCron.GetTimeForType(checkTimeType, s.Torrents[hash])
					// 判断时间类型
					if tmpTime > minScanTime {
						hashesChecked = append(hashesChecked, hash)
					}

				}
			}

			//>> 通知设置分享率
			if len(hashesChecked) > 0 {
				// 分段处理避免过多
				list := service.ServiceHelper.ArraySplit(hashesChecked, 6)

				for _, v := range list {
					log.Println("种子已加入分享限制:")
					for _, v2 := range v {
						log.Println(s.Torrents[v2].Name)
					}
					// 判断是否需要设置分享速率
					if uploadLimit > 0 {
						// 通知设置分享率
						err, _ := service.ServiceCron.GetTorrents().SetUploadLimit(torrents.ApiTorrentSetUploadLimitReq{
							Hashes: strings.Join(v, "|"),
							Limit:  uploadLimit,
						})
						if err != nil {
							log.Println("设置上传速度失败", err.Error())
						} else {
							log.Println("设置上传速度成功 共:", len(v))
						}
					}

					// 通知设置分享率
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
			this.checkScanTime(len(hashes))
			// 检查登录状态是否合法
			this.checkLogin()

			log.Println("======================================\r\n\r\n")

		}
	}()
	for {
		// 外部监听避免掉线
		time.Sleep(time.Second * 10)
	}
}

// checkScanTime 检查扫描时间间隔
func (this *handleCron) checkScanTime(setLimitCount int) {
	limitTime := this.conf.GetDuration("qbit_scan_time")
	if limitTime == 0 {
		this.conf.Set("qbit_scan_time", "10")
		err := this.conf.WriteConfig()
		if err != nil {
			log.Println("写入配置失败,请注意配置文件权限")
		}
		limitTime = 10
	}
	if setLimitCount > 0 {
		log.Println("扫描完成 \r\n\t\t共扫描到提交限制分享种子数量:" + strconv.Itoa(setLimitCount) + "个")
	}
	log.Println("开始等待下一轮检查 等待", int(limitTime), "s后执行")
	time.Sleep(time.Second * limitTime)
}

func (this *handleCron) checkLogin() {
gotoCheckLogin:
	if service.ServiceCron.CheckLogin() == false {
		log.Println("登录失败,等待10分钟后重试")
		time.Sleep(time.Second * 60 * 10)
		goto gotoCheckLogin
	}
}
