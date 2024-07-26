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

### Docker
docker 地址:
https://hub.docker.com/r/xiaoyi510/qbit-auto-limit

docker-compose.yml
```
version: "3.3"
services:
  qbit-auto-limit:
    container_name: qbit-auto-limit
    environment:
      - TZ=Asia/Shanghai
      - HOST_OS=Unraid
      - HOST_HOSTNAME=Tower
      - HOST_CONTAINERNAME=qbit-auto-limit-new
    labels:
      - net.unraid.docker.managed=dockerman
    volumes:
      - /mnt/user/appdata/qbit-auto-limit/conf:/app/conf
    image: xiaoyi510/qbit-auto-limit:latest
networks: {}
```

/mnt/user/appdata/qbit-auto-limit/conf 为配置目录
