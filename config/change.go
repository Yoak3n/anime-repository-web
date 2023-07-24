package config

import "net/url"

func ChangeAllConfig(key string, use bool, proxy string, rawPath string, tVPath string, moviePath string, delay int64) error {
	Conf.ApiKey = key
	Conf.UseProxy = use
	parse, err := url.Parse(proxy)
	if err != nil {
		return err
	}
	Conf.Proxy = parse
	Conf.RawPath = rawPath
	Conf.TVPath = tVPath
	Conf.MoviePath = moviePath
	Conf.Delay = delay

	v.Set("api_key", key)
	v.Set("use_proxy", use)
	v.Set("proxy", proxy)
	v.Set("raw_path", rawPath)
	v.Set("tv_path", tVPath)
	v.Set("movie_path", moviePath)
	v.Set("delay", delay)
	err = v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeAPIKey(key string) error {
	Conf.ApiKey = key
	v.Set("api_key", key)
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeProxy(use bool, proxy string) error {
	if use {
		parse, err := url.Parse(proxy)
		if err != nil {
			return err
		}
		Conf.UseProxy = true
		Conf.Proxy = parse
	} else {
		Conf.UseProxy = false
		Conf.Proxy = nil
	}
	v.Set("use_proxy", use)
	v.Set("proxy", proxy)
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeRawPath(path string) error {
	Conf.RawPath = path
	v.Set("raw_path", path)
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeTVPath(path string) error {
	Conf.TVPath = path
	v.Set("tv_path", path)
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeMoviePath(path string) error {
	Conf.MoviePath = path
	v.Set("movie_path", path)
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}

func ChangeDelay(delay int64) error {
	Conf.Delay = delay
	v.Set("delay", delay)
	err := v.WriteConfig()
	if err != nil {
		return err
	}
	return nil
}
