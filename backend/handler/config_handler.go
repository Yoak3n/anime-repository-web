package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github/Yoak3n/anime-repository-web/backend/backthread"
	"github/Yoak3n/anime-repository-web/config"
	"github/Yoak3n/anime-repository-web/package/response"
	"net"
	"net/http"
	"os"
)

func GetConfig(c *gin.Context) {
	response.SuccessWithData(c, response.ExposeConfig(config.Conf))
}

func GetHost(c *gin.Context) {
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
	response.SuccessWithData(c, hostname)
}
func ChangeConfig(c *gin.Context) {
	rawJson, err := c.GetRawData()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, "Get json error", nil)
	} else {
		result := gjson.ParseBytes(rawJson)
		key := result.Get("key").String()
		useProxy := result.Get("use_proxy").Bool()
		proxyUrl := result.Get("proxy_url").String()
		rawPath := result.Get("raw_path").String()
		tvPath := result.Get("tv_path").String()
		moviePath := result.Get("movie_path").String()
		delay := result.Get("delay").Int()
		err = config.ChangeAllConfig(key, useProxy, proxyUrl, rawPath, tvPath, moviePath, delay)
		if err != nil {
			response.Response(c, http.StatusInternalServerError, "Change config error", nil)
		} else {
			_ = backthread.Scan()
			response.Success(c)
		}
	}
}
