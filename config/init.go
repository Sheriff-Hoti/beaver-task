package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/Sheriff-Hoti/beaver-task/data"
)

type Config struct {
	Test    string `json:"test"`
	DataDir string `json:"data_dir"`
}

// func (*Config) Validate() error {
// 	if
// }

func GetDefaultConfigVals() *Config {

	return &Config{
		Test:    "testing",
		DataDir: data.GetDefaultDataPath(),
	}
}

func GetDefaultConfigPath() string {
	xdgConfig := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfig == "" {
		xdgConfig = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return filepath.Join(xdgConfig, "beaver-task", "config.json")
}

func ReadConfigFile(config_path string) (*Config, error) {

	if _, err := os.Stat(config_path); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist and if it does not exists just return the defaults

		return GetDefaultConfigVals(), nil
	}

	file, err := os.Open(config_path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var config Config

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
