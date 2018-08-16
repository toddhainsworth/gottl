package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/jason0x43/go-toggl"
)

var workspaces = flag.Bool("workspaces", false, "get workspaces and their IDs")

func main() {
	toggl.DisableLog()
	flag.Parse()
	config, err := NewConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	if *workspaces {
		app := App{APIKey: config.APIKey}
		app.StartSession()
		app.PrintWorkspaces()
		os.Exit(0)
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
