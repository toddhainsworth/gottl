package main

import (
	"flag"
	"fmt"
	"os"

	toggl "github.com/machiel/go-toggl"
)

var workspaces = flag.Bool("workspaces", false, "get workspaces and their IDs")
var current = flag.Bool("current", false, "get the currently running timer (if available)")

func main() {
	// TODO: configurable
	toggl.AppName = "Gottl Toggl Util"

	toggl.DisableLog()

	flag.Parse()
	config, err := NewConfig()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *workspaces {
		app := App{APIKey: config.APIKey}
		app.StartSession()
		app.PrintWorkspaces()
		os.Exit(0)
	}

	if config.Workspace <= 0 {
		fmt.Println("No workspace configured, run again with --workspaces to find your workspace ID")
		os.Exit(1)
	}

	app := App{APIKey: config.APIKey, WorkspaceID: config.Workspace}
	err = app.StartSession()

	if err != nil {
		fmt.Println(err)
	}

	if *current {
		err = app.PrintCurrentTimer()
	} else {
		err = app.PrintReport()
	}

	if err != nil {
		fmt.Println(err)
	}
}
