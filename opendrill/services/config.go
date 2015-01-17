package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Config gathers all the necessary data to run the app
type Config struct {
	URL          string `json:"url"`
	Port         string `json:"port"`
	DatabaseUrl  string `json:"database_url"`
	DatabaseName string `json:"database_name"`
	Debug        bool   `json:"debug"`
}

// NewConfig creates a new config struct
func NewConfig(configPath string) (*Config, error) {
	var config = new(Config)

	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
