package router

import (
	"github.com/gin-gonic/gin"
	"github/Yoak3n/anime-repository-web/backend/handler"
	"net/http"
)

func renderIndex(c *gin.Context) {
	r.LoadHTMLFiles("resource/index.html")
	//r.StaticFile("favicon.ico", "resource/html/favicon.ico")
	c.HTML(http.StatusOK, "index.html", "")
}

func runController() {
	r.GET("/", renderIndex)
	r.Static("/assets", "resource/assets")
	r.StaticFile("/vite.svg", "resource/vite.svg")
	apiV1()
}

func apiV1() {
	v1 := r.Group("/v1")

	v1.GET("/host", handler.GetHost)
	// ? 一个问号表示我对这接口的疑问，开发初期留下来的垃圾
	v1.POST("/tv/:id/:language", handler.GetTvInfo)
	apiV1Config(v1)
	apiV1Rule(v1)
	apiV1Path(v1)
	apiV1Raw(v1)
}

func apiV1Config(v1 *gin.RouterGroup) {
	config := v1.Group("/config")

	config.GET("", handler.GetConfig)
	config.POST("", handler.ChangeConfig)
}

func apiV1Path(v1 *gin.RouterGroup) {
	path := v1.Group("/path")

	path.GET("/", handler.GetPaths)
	path.POST("/raw/:path", handler.ChangeRawPath)
	path.POST("/tv/:path", handler.ChangeTVPath)
	path.POST("/movie/:path", handler.ChangeMoviePath)
}

func apiV1Rule(v1 *gin.RouterGroup) {
	rule := v1.Group("/rule")

	rule.GET("", handler.GetRule)
	rule.POST("", handler.AddRule)
	rule.PATCH("/:id", handler.UpdateRule)
	rule.DELETE("/:id", handler.HiddenRule)
}

func apiV1Raw(v1 *gin.RouterGroup) {
	raw := v1.Group("/raw")

	raw.GET("", handler.GetRaw)
	raw.GET("/tv/:vid", handler.GetRawTVName)
}
