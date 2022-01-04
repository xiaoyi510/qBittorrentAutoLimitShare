# qBittorrentAutoLimitShare

qBittorrent Web Api Auto Set Upload Limit

### 配置项
trust_trackers 信任tracker 一行一个

qbit_cookie qbit cookie

qbit_upload_radio 不信任tracker分享率设置

qbit_upload_time 不信任tracker对应种子分享上传设置


### API 完成度

| 模块          | 接口                                | 完成  |
|-------------|-----------------------------------|-----|
| Auth        | auth/login                        | ✔️  |
|             | auth/logout                       | ✔️  |
| -           | -                                 | -️  |
| Application |                                   | ✔️  |
|             | app/version                       | ✔️  |
|             | app/webapiVersion                 | ✔️  |
|             | app/buildInfo                     | ✔️  |
|             | app/shutdown                      | ✔️  |
|             | app/preferences                   | ✔️  |
|             | app/setPreferences                | ✔️  |
|             | app/defaultSavePath               | ✔️  |
| Log         |                                   | ✔️  |
|             | log/main                          | ✔️  |
|             | log/peers                         | ✔️  |
| Sync        |                                   | ✔️  |
|             | sync/maindata                     | ✔️  |
|             | sync/torrentPeers                 | ✔️  |
| transfer    |                                   | ❌️  |
|             | transfer/info                     | ❌️  |
|             | transfer/speedLimitsMode          | ❌️  |
|             | transfer/toggleSpeedLimitsMode    | ❌️  |
|             | transfer/downloadLimit            | ❌️  |
|             | transfer/setDownloadLimit         | ❌️  |
|             | transfer/uploadLimit              | ❌️  |
|             | transfer/setUploadLimit           | ❌️  |
|             | transfer/banPeers                 | ❌️  |
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