package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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
		log.Printf("Error reading config file (%v), falling back to environment variables", err)
	} else {
		err = yaml.Unmarshal(file, config)
		if err != nil {
			log.Printf("Error unmarshalling YAML from file (%v), falling back to environment variables", err)
		}

		return config, nil
	}

	// fallback to env vars if config file not available
	config.Database.User = os.Getenv("DATABASE_USER")
	config.Database.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database.Host = os.Getenv("DATABASE_HOST")
	config.Database.Sslmode = os.Getenv("DATABASE_SSLMODE")
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalf("error convert to int: %v", err)
		return nil, err
	}
	config.Database.Port = port
	config.Database.Dbname = os.Getenv("DATABASE_NAME")

	if config.Database.User == "" || config.Database.Password == "" || config.Database.Host == "" || config.Database.Dbname == "" {
		log.Fatalf("Missing required configuration for database connection from environment variables")
		return nil, fmt.Errorf("missing required environment variables")
	}

	return config, nil
}
