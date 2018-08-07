package main

import (
	"fmt"

	"github.com/jason0x43/go-toggl"
)

func main() {
	toggl.DisableLog()
	config, err := NewConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	app := App{APIKey: config.APIKey, WorkspaceID: config.Workspace}
	err = app.StartSession()

	if err != nil {
		fmt.Println(err)
	}

	err = app.PrintReport()

	if err != nil {
		fmt.Println(err)
	}
}
