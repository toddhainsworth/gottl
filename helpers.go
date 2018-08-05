package main

import (
	"io/ioutil"
	"os"
	"strings"
)

// GetAPIKey returns the API key stored in the configuration file
func GetAPIKey() (string, error) {
	data, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.gottl"))

	if err != nil {
		return "", err
	}

	str := string(data)
	return strings.TrimSpace(str), nil
}
