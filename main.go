package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jason0x43/go-toggl"
)

func main() {
	toggl.DisableLog()
	apiKey, err := getAPIKey()

	if err != nil {
		fmt.Println(err)
		return
	}

	app := App{APIKey: apiKey}
	err = app.StartSession()

	if err != nil {
		fmt.Println(err)
	}

	err = app.PrintReport()

	if err != nil {
		fmt.Println(err)
	}
}

// Get the API key from the .gottl file
func getAPIKey() (string, error) {
	data, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.gottl"))

	if err != nil {
		return "", err
	}

	str := string(data)
	return strings.TrimSpace(str), nil
}
