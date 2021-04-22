package link_shorter

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	MySqlDB struct {
		Host     string `yaml:"host"`
		DbName   string `yaml:"db_name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
	Server struct {
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

func NewConfig() *Config {
	cfg := &Config{}
	cfg.MySqlDB.Host = "127.0.0.1"
	cfg.MySqlDB.DbName = "testdb"
	cfg.MySqlDB.User = "root"
	cfg.MySqlDB.Password = "root"

	cfg.Server.Port = "55081"
	cfg.Log.Level = "trace"
	return cfg
}
