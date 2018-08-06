package main

import (
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

const configPath = "$HOME/.gottl"

// Config represents the Gottl config
type Config struct {
	APIKey    string `yaml:"api_key"`
	Workspace int    `yaml:"workspace"`
}

// NewConfig loads and reutrns a Config struct
func NewConfig() (Config, error) {
	config := Config{}
	data, err := ioutil.ReadFile(os.ExpandEnv(configPath))

	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(data), &config)

	if err != nil {
		return config, err
	}

	return config, nil
}
