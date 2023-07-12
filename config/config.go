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
	Port      int      `yaml:"port"`
	ApiKey    string   `yaml:"apiKey"`
	Proxy     *url.URL `yaml:"proxy"`
	RawPath   string   `yaml:"raw_path"`
	TVPath    string   `yaml:"tv_path"`
	MoviePath string   `yaml:"movie_path"`
	Delay     int64    `yaml:"delay"`
}

func init() {
	v = viper.New()
	Conf = new(Configuration)
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	err := v.ReadInConfig()
	if err != nil {
		log.Panic("请检查配置文件", err)
	}
	Conf.Port = v.GetInt("port")
	Conf.ApiKey = v.GetString("apiKey")
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
	v.WatchConfig()
}

func (c *Configuration) handlerPath() {
	c.TVPath = util.ReplacePressDown(c.TVPath)
	c.RawPath = util.ReplacePressDown(c.RawPath)
	c.MoviePath = util.ReplacePressDown(c.MoviePath)
}
