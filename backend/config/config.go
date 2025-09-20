package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Dbname   string `yaml:"dbname"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"database"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("error unmarshalling yaml: %v", err)
		return nil, err
	}

	return config, nil
}
