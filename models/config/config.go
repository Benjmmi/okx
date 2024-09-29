package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Demo         bool
	Colo         bool
	SourceIPs    []string
	TargetIPs    []string
	OkxAPIKey    string
	OkxSecretKey string
	OkxPassword  string
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config Config
	if err = json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
