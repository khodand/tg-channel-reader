package config

import (
	"os"

	"encoding/json"
	"github.com/pkg/errors"
)

const FilePath = "secret_config.json"

type Config struct {
	Token          string `json:"token"`
	Payment        string `json:"payment"`
	UpdatesThreads int    `json:"updates_threads"`
	CommonThreads  int    `json:"common_threads"`
	PremiumThreads int    `json:"premium_threads"`
}

func InitConfig() (*Config, error) {
	var config Config
	jsonFile, err := os.ReadFile(FilePath)
	if err != nil {
		return nil, errors.Wrap(err, "read failed")
	}
	err = json.Unmarshal(jsonFile, &config)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal failed")
	}

	return &config, nil
}
