# qBittorrentAutoLimitShare

可用于PT下载和其他非PT站下载管理分享率

一般情况下PT要保种 不限制分享率

但是其他BT下载网站则需要限制分享率= =不过你的资源够也无所谓

### 配置项

```yaml
# 限制上传速度 0 不限制
qbit_upload_limit: "1"
# 限制上传分享率
qbit_upload_radio: "1.0"
# 限制上传分享时间
qbit_upload_time: "-1"
# 检测跳过多少天以前的种子
qbit_skip_max_complete_time: "1"
# 检测时间类型 1 活动时间 2 添加时间 3完成时间
qbit_check_time_type: "1"
# 扫描间隔时间单位s
qbit_scan_time: "10"
# qBit服务器设置
qbit_server:
  # 可以直接填写cookie SID=xx;
  cookie: ""
  # qBit服务器域名/IP
  url: ""
  # qBit WebUI端口号
  port: ""
  # qBit 是否为ssl
  ssl: "1"
  # qBit 账号密码
  username: ""
  password: ""
# 信任的tracker 不处理限制上传分享率
trust_trackers: "tracker.hdtime.org/announce.php
hdfans.org/announce.php
"
# 超过这个tracker数量的自动限制分享率 为0 不处理
tracker_max: "0"
```

### 版本更新

v0.0.2

```yaml
# 超过这个tracker数量的自动限制分享率 为0 不处理
tracker_max: "0"
```

#### v0.0.1

新增配置项

```yaml
#检测时间类型 1 活动时间 2 添加时间 3完成时间
qbit_check_time_type: "1"
```

### API 完成度


| 模块        | 接口                              | 完成 |
| ----------- | --------------------------------- | ---- |
| Auth        | auth/login                        | ✔️ |
|             | auth/logout                       | ✔️ |
| -           | -                                 | -️  |
| Application |                                   | ✔️ |
|             | app/version                       | ✔️ |
|             | app/webapiVersion                 | ✔️ |
|             | app/buildInfo                     | ✔️ |
|             | app/shutdown                      | ✔️ |
|             | app/preferences                   | ✔️ |
|             | app/setPreferences                | ✔️ |
|             | app/defaultSavePath               | ✔️ |
| Log         |                                   | ✔️ |
|             | log/main                          | ✔️ |
|             | log/peers                         | ✔️ |
| Sync        |                                   | ✔️ |
|             | sync/maindata                     | ✔️ |
|             | sync/torrentPeers                 | ✔️ |
| transfer    |                                   | ❌️ |
|             | transfer/info                     | ❌️ |
|             | transfer/speedLimitsMode          | ❌️ |
|             | transfer/toggleSpeedLimitsMode    | ❌️ |
|             | transfer/downloadLimit            | ❌️ |
|             | transfer/setDownloadLimit         | ❌️ |
|             | transfer/uploadLimit              | ❌️ |
|             | transfer/setUploadLimit           | ❌️ |
|             | transfer/banPeers                 | ❌️ |
| torrents    |                                   | ❌   |
|             | torrents/info                     | ❌   |
|             | torrents/properties               | ❌   |
|             | torrents/trackers                 | ❌   |
|             | torrents/webseeds                 | ❌   |
|             | torrents/files                    | ❌   |
|             | torrents/pieceStates              | ❌   |
|             | torrents/pieceHashes              | ❌   |
|             | torrents/pause                    | ❌   |
|             | torrents/resume                   | ❌   |
|             | torrents/delete                   | ❌   |
|             | torrents/recheck                  | ❌   |
|             | torrents/reannounce               | ❌   |
|             | torrents/addTrackers              | ❌   |
|             | torrents/add                      | ❌   |
|             | torrents/editTracker              | ❌   |
|             | torrents/removeTrackers           | ❌   |
|             | torrents/addPeers                 | ❌   |
|             | torrents/increasePrio             | ❌   |
|             | torrents/decreasePrio             | ❌   |
|             | torrents/topPrio                  | ❌   |
|             | torrents/bottomPrio               | ❌   |
|             | torrents/filePrio                 | ❌   |
|             | torrents/downloadLimit            | ❌   |
|             | torrents/setDownloadLimit         | ❌   |
|             | torrents/setShareLimits           | ❌   |
|             | torrents/uploadLimit              | ❌   |
|             | torrents/setUploadLimit           | ❌   |
|             | torrents/setLocation              | ❌   |
|             | torrents/rename                   | ❌   |
|             | torrents/setCategory              | ❌   |
|             | torrents/categories               | ❌   |
|             | torrents/createCategory           | ❌   |
|             | torrents/editCategory             | ❌   |
|             | torrents/removeCategories         | ❌   |
|             | torrents/addTags                  | ❌   |
|             | torrents/removeTags               | ❌   |
|             | torrents/tags                     | ❌   |
|             | torrents/createTags               | ❌   |
|             | torrents/deleteTags               | ❌   |
|             | torrents/setAutoManagement        | ❌   |
|             | torrents/toggleSequentialDownload | ❌   |
|             | torrents/toggleFirstLastPiecePrio | ❌   |
|             | torrents/setForceStart            | ❌   |
|             | torrents/setSuperSeeding          | ❌   |
|             | torrents/renameFile               | ❌   |
|             | torrents/renameFolder             | ❌   |
| rss         |                                   | ❌   |
|             | rss/addFolder                     | ❌   |
|             | rss/addFeed                       | ❌   |
|             | rss/removeItem                    | ❌   |
|             | rss/moveItem                      | ❌   |
|             | rss/items                         | ❌   |
|             | rss/markAsRead                    | ❌   |
|             | rss/refreshItem                   | ❌   |
|             | rss/setRule                       | ❌   |
|             | rss/renameRule                    | ❌   |
|             | rss/removeRule                    | ❌   |
|             | rss/rules                         | ❌   |
| search      |                                   | ❌   |
|             | search/start                      | ❌   |
|             | search/stop                       | ❌   |
|             | search/status                     | ❌   |
|             | search/results                    | ❌   |
|             | search/delete                     | ❌   |
|             | search/plugins                    | ❌   |
|             | search/installPlugin              | ❌   |
|             | search/enablePlugin               | ❌   |
|             | search/updatePlugins              | ❌   |
