package request

import (
	"github/Yoak3n/anime-repository-web/config"
	"net/http"
	"time"
)

func GenRequestClient() http.Client {
	if config.Conf.UseProxy {
		return genRequestClientWithProxy()
	} else {
		return genRequestClientWithoutProxy()
	}
}

func genRequestClientWithProxy() http.Client {
	return http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 15 * time.Second,
			Proxy:           http.ProxyURL(config.Conf.Proxy),
		},
	}
}
func genRequestClientWithoutProxy() http.Client {
	return http.Client{
		Timeout: time.Second * 5,
	}
}
