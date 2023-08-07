package router

import (
	"github.com/gin-gonic/gin"
	"github/Yoak3n/anime-repository-web/backend/handler"
	"net"
	"net/http"
	"os"
)

func renderIndex(c *gin.Context) {
	r.LoadHTMLFiles("resource/index.html")
	//r.StaticFile("favicon.ico", "resource/html/favicon.ico")

	hostname, err := os.Hostname()
	if err != nil {
		ifaces, err := net.Interfaces()
		if err != nil {
			hostname = "unknown"
		}
		for _, iface := range ifaces {
			if iface.Flags&net.FlagUp == 0 {
				continue
			}
			if iface.Flags&net.FlagLoopback != 0 {
				continue
			}
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ip == nil || ip.IsLoopback() {
					continue
				}
				ip = ip.To4()
				if ip == nil {
					continue
				} else {
					hostname = ip.String()
				}
			}
		}
	}
	c.HTML(http.StatusOK, "index.html", hostname)
}

func runController() {
	r.GET("/", renderIndex)
	r.Static("/assets", "resource/assets")
	r.StaticFile("/vite.svg", "resource/vite.svg")
	apiV1()
}

func apiV1() {
	v1 := r.Group("/v1")

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
