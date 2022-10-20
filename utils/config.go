package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	ApiKey string `yaml:"api_key_bot"`
}

func GetConfig() config {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	var cfg config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
	return cfg
}

func processError(err error) {
	panic("unimplemented")
}
