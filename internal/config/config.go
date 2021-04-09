package config

import (
	"io/ioutil"
)

type Config struct {
	kubeConfigs []string
}

func ReadConfig(configFile string) (string, error) {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
