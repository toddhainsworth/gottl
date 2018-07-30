package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/jason0x43/go-toggl"
)

func main() {
	toggl.DisableLog()
	apiKey, err := getAPIKey()

	if err != nil {
		fmt.Println(err)
		return
	}

	session := toggl.OpenSession(apiKey)
	account, err := session.GetAccount()

	if err != nil {
		fmt.Println(err)
		return
	}

	start, end := getDates()
	for _, workspace := range account.Data.Workspaces {
		// Only get the first page, any more data and we'd likely fill the terminal
		report, err := session.GetDetailedReport(workspace.ID, start, end, 1)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Report for \"%s\"\n", workspace.Name)

		itemsByTime := make(map[string][]toggl.DetailedTimeEntry, len(report.Data))
		for _, reportItem := range report.Data {
			startString := getStartString(*reportItem.Start)
			// If the map has the time key already then we can append, otherwise insert
			if _, ok := itemsByTime[startString]; ok {
				itemsByTime[startString] = append(itemsByTime[startString], reportItem)
			} else {
				itemsByTime[startString] = []toggl.DetailedTimeEntry{reportItem}
			}
			// printItem(reportItem)
		}

		for time, items := range itemsByTime {
			fmt.Println(time)
			for _, item := range items {
				printItem(item)
			}
		}
	}
}

func getStartString(time time.Time) string {
	return fmt.Sprintf("-- %s %d/%d/%d --",
		time.Weekday(), time.Day(), time.Month(), time.Year())
}

// Print the given item in the regular format
func printItem(item toggl.DetailedTimeEntry) {
	dur, err := getDuration(item.Duration)

	if err != nil {
		fmt.Println(err)
		return
	}

	startTime, endTime := item.Start.Format("01:02"), item.End.Format("01:02")
	fmt.Printf("(%s - %s) %s - %s\n",
		startTime, endTime, item.Description, dur.String())
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

// Converts the duration in 'ms' to long-form
func getDuration(ms int64) (time.Duration, error) {
	return time.ParseDuration(fmt.Sprintf("%dms", ms))
}

// Helper function to pretty-print the struct
func ppStruct(data interface{}) {
	fmt.Printf("%+v\n", data)
}
