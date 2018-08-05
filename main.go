package main

import (
	"fmt"

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
