package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jason0x43/go-toggl"
)

// App is the main app structure
type App struct {
	APIKey    string
	Workspace int
	session   toggl.Session
}

// NewApp creates an App struct
func NewApp(APIKey string) App {
	return App{APIKey: APIKey}
}

// StartSession starts a Toggl session and stores it on the App
func (app *App) StartSession() error {
	if app.APIKey == "" {
		return errors.New("API key is not present")
	}

	app.session = toggl.OpenSession(app.APIKey)
	return nil
}

// PrintReport fetches time entries and constructs a report for the CLI
func (app App) PrintReport() error {
	if app.session.APIToken != app.APIKey {
		return errors.New("Session is not active")
	}

	start, end := getDates()
	report, err := app.session.GetDetailedReport(app.Workspace, start, end, 1)

	if err != nil {
		return err
	}

	fmt.Println(report)
	itemsByTime := getItemsByTime(report)

	for time, items := range itemsByTime {
		fmt.Println(time)
		for _, item := range items {
			if err := printItem(item); err != nil {
				return err
			}
		}
	}

	return nil
}

// Group report items by their start time
func getItemsByTime(report toggl.DetailedReport) map[string][]toggl.DetailedTimeEntry {
	itemsByTime := make(map[string][]toggl.DetailedTimeEntry, len(report.Data))
	for _, reportItem := range report.Data {
		startString := getStartString(*reportItem.Start)
		// If the map has the time key already then we can append, otherwise insert
		if _, ok := itemsByTime[startString]; ok {
			itemsByTime[startString] = append(itemsByTime[startString], reportItem)
		} else {
			itemsByTime[startString] = []toggl.DetailedTimeEntry{reportItem}
		}
	}

	return itemsByTime
}

func (app App) getAccount() (toggl.Account, error) {
	return app.session.GetAccount()
}

// Get the start and end dates to send to the Toggl API
func getDates() (start, end string) {
	endDate := time.Now()
	y, m, d := endDate.Date()
	end = fmt.Sprintf("%d-%d-%d", y, m, d)

	// Subtract 5 days
	startDate := endDate.AddDate(0, 0, -5)
	y, m, d = startDate.Date()
	start = fmt.Sprintf("%d-%d-%d", y, m, d)
	return

}

func getStartString(time time.Time) string {
	return fmt.Sprintf("-- %s %d/%d/%d --",
		time.Weekday(), time.Day(), time.Month(), time.Year())
}

// Print the given item in the regular format
func printItem(item toggl.DetailedTimeEntry) error {
	dur, err := getDuration(item.Duration)

	if err != nil {
		return err
	}

	startTime, endTime := item.Start.Format("15:04"), item.End.Format("15:04")
	fmt.Printf("(%s - %s) %s - %s\n",
		startTime, endTime, item.Description, dur.String())
	return nil
}

// Converts the duration in 'ms' to long-form
func getDuration(ms int64) (time.Duration, error) {
	return time.ParseDuration(fmt.Sprintf("%dms", ms))
}
