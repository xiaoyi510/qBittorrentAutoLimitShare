package handler

import (
	"QbittorrentAutoLimitShare/internal/model/http"
	"QbittorrentAutoLimitShare/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

var HandlerOptions = &options{}

type options struct {
}

func (this *options) SetConf(c *gin.Context) {
	// 保存配置项
	var req http.OptionsSetReq
	err := c.ShouldBind(&req)
	if err != nil {
		log.Println(err)
		return
	}

	service.ServiceHttpOptions.Set(req)
}
