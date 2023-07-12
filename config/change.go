package config

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
