package config

import (
	"github.com/spf13/viper"
	"github/Yoak3n/anime-repository-web/package/util"
	"log"
	"net/url"
)

var (
	v    *viper.Viper
	Conf *Configuration
)

type Configuration struct {
	Port      int      `yaml:"port" json:"port"`
	ApiKey    string   `yaml:"api_key" json:"api_key"`
	UseProxy  bool     `yaml:"use_proxy" json:"use_proxy"`
	Proxy     *url.URL `yaml:"proxy"  json:"proxy"`
	RawPath   string   `yaml:"raw_path" json:"raw_path"`
	TVPath    string   `yaml:"tv_path" json:"tv_path"`
	MoviePath string   `yaml:"movie_path" json:"movie_path"`
	Delay     int64    `yaml:"delay" json:"delay"`
}

func init() {
	v = viper.New()
	Conf = new(Configuration)

	v.SetDefault("port", 8080)
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	err := v.ReadInConfig()
	if err != nil {
		log.Panic("请检查配置文件", err)
	}
	Conf.Port = v.GetInt("port")
	Conf.ApiKey = v.GetString("api_key")
	Conf.UseProxy = v.GetBool("use_proxy")
	parse, err := url.Parse(v.GetString("proxy"))
	if err != nil {
		log.Panic("请检查配置文件", err)
	}
	Conf.Proxy = parse

	Conf.TVPath = v.GetString("tv_path")
	Conf.MoviePath = v.GetString("movie_path")
	Conf.Delay = v.GetInt64("delay")
	Conf.RawPath = v.GetString("raw_path")
	Conf.handlerPath()
	Conf.setViperKey()
	v.WatchConfig()
}

func (c *Configuration) handlerPath() {
	c.TVPath = util.ReplacePressDown(c.TVPath)
	c.RawPath = util.ReplacePressDown(c.RawPath)
	c.MoviePath = util.ReplacePressDown(c.MoviePath)
}
func (c *Configuration) setViperKey() {
	v.Set("port", c.Port)
	v.Set("api_key", c.ApiKey)
	v.Set("proxy", c.Proxy)
	v.Set("tv_path", c.TVPath)
	v.Set("movie_path", c.MoviePath)
	v.Set("delay", c.Delay)
	v.Set("raw_path", c.RawPath)
}
