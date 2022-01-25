package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var PublicWeb = &publicWeb{}

type publicWeb struct {
	engine *gin.Engine
}

// Run 运行http服务
func (this *publicWeb) Run() {
	this.engine = gin.Default()

	// 加载html文件
	this.engine.LoadHTMLGlob("public/*")
	// 注册路由
	this.regRouter()

	go this.engine.Run("0.0.0.0:8080")
}

func (this *publicWeb) regRouter() {
	index := func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{"title": "设置"})
	}

	this.engine.GET("/", index)
	this.engine.GET("/index", index)
	this.engine.POST("/set-conf", HandlerOptions.SetConf)

}
