package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	toggl "github.com/machiel/go-toggl"
)

// App is the main app structure
type App struct {
	APIKey      string
	WorkspaceID int
	session     toggl.Session
}

// NewApp creates an App struct
func NewApp(apiKey string, workspaceID int) App {
	return App{APIKey: apiKey, WorkspaceID: workspaceID}
}

// StartSession starts a Toggl session and stores it on the App
func (app *App) StartSession() error {
	if app.APIKey == "" {
		return errors.New("API key is not present")
	}

	app.session = toggl.OpenSession(app.APIKey)
	return nil
}

// PrintCurrentTimer prints the current running timer (if available) to STDOUT
func (app App) PrintCurrentTimer() error {
	if app.session.APIToken != app.APIKey {
		return errors.New("Session is not active")
	}

	entry, err := app.session.GetCurrentTimeEntry()

	if !entry.IsRunning() {
		return errors.New("No current timer")
	}

	if err != nil {
		return err
	}

	return printTimeEntry(entry)
}

// PrintReport fetches time entries and constructs a report for the CLI
func (app App) PrintReport() error {
	if app.session.APIToken != app.APIKey {
		return errors.New("Session is not active")
	}

	start, end := getDates()
	workspace, err := app.getWorkspace()

	if err != nil {
		return err
	}

	report, err := app.session.GetDetailedReport(workspace.ID, start, end, 1)

	if err != nil {
		return err
	}

	dates, itemsByTime := getItemsByTime(report)

	for _, date := range dates {
		items := itemsByTime[date]

		color.Green(date)
		// Holding a daily tally
		var dayDuration int64

		for _, item := range items {
			if err := printDetailedTimeEntry(item); err != nil {
				return err
			}
			dayDuration += item.Duration
		}

		duration, err := getDuration(dayDuration)
		if err != nil {
			return err
		}
		color.Magenta("Total: %s", duration)
	}

	duration, err := getDuration(int64(report.TotalGrand))
	if err != nil {
		return err
	}
	color.Green("Grand Total: %s", duration)

	return nil
}

func (app App) getAccount() (toggl.Account, error) {
	return app.session.GetAccount()
}

// PrintWorkspaces gets a list of all workspaces such that the user
// can then use their ID in config
func (app App) PrintWorkspaces() error {
	if app.session.APIToken != app.APIKey {
		return errors.New("Session is not active")
	}

	account, err := app.getAccount()

	if err != nil {
		return err
	}

	for _, workspace := range account.Data.Workspaces {
		fmt.Printf("Workspace #%d: %s\n", workspace.ID, workspace.Name)
	}

	return nil
}

func (app *App) getWorkspace() (toggl.Workspace, error) {
	account, err := app.getAccount()

	if err != nil {
		return toggl.Workspace{}, err
	}

	for _, workspace := range account.Data.Workspaces {
		if workspace.ID == app.WorkspaceID {
			return workspace, nil
		}
	}

	return toggl.Workspace{}, nil
}
