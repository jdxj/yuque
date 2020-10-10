package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	cfg *configuration
)

func Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	cfg = new(configuration)
	return decoder.Decode(cfg)
}

type configuration struct {
	Token string `yaml:"token"`
}

func Token() string {
	return cfg.Token
}
