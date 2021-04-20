package link_shorter

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	MySqlDB struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`
}

func ReadConf(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}

	if err := yaml.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
