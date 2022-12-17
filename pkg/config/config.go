package config

import (
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

var HOME string = os.Getenv("HOME")
var FILEPATH string = path.Join(HOME, ".config/outline")

type Config struct {
	Host  string `yaml:"host,omitempty"`
	Token string `yaml:"token,omitempty"`
}

func (h *Config) Write() error {
	data, err := yaml.Marshal(h)
	if err != nil {
		return err
	}
	// log.Printf("Creating config: %s", FILEPATH)
	err = os.MkdirAll(FILEPATH, 0733)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("%s/config.yaml", FILEPATH), data, 0733)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Read() error {
	file, err := os.ReadFile(fmt.Sprintf("%s/config.yaml", FILEPATH))
	if err != nil {
		return fmt.Errorf("coulf not read file: %v", err)
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		fmt.Printf("File: could read config file: %v", err)
		return fmt.Errorf("could read config file: %v", err)
	}
	return nil
}
