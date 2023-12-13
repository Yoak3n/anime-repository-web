package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github/Yoak3n/anime-repository-web/config"
	"log"
)

var r *gin.Engine

func RunSever() {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	runController()
	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		log.Panic(fmt.Sprintf("服务端启动失败：%v", err))
	}
}
