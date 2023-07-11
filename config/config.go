package config

import (
	"github.com/spf13/viper"
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
	Conf.TVPath = v.GetString("tv_path")
	if err != nil {
		log.Panic("请检查配置文件", err)
	}
	Conf.Proxy = parse
	v.WatchConfig()
}

func ChangeRawPath(path string) error {
	Conf.RawPath = path
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeTVPath(path string) error {
	Conf.TVPath = path
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeMoviePath(path string) error {
	Conf.MoviePath = path
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
