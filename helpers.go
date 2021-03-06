package main

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	toggl "github.com/machiel/go-toggl"
)

// Group report items by their start time
func getItemsByTime(report toggl.DetailedReport) ([]string, map[string][]toggl.DetailedTimeEntry) {
	itemsByTime := make(map[string][]toggl.DetailedTimeEntry, len(report.Data))
	keys := []string{}

	for _, reportItem := range report.Data {
		startString := getStartString(*reportItem.Start)
		// If the map has the time key already then we can append, otherwise insert
		if _, ok := itemsByTime[startString]; ok {
			itemsByTime[startString] = append(itemsByTime[startString], reportItem)
		} else {
			itemsByTime[startString] = []toggl.DetailedTimeEntry{reportItem}
			keys = append(keys, startString)
		}
	}

	return keys, itemsByTime
}

// Get the start and end dates to send to the Toggl API
func getDates() (start, end string) {
	endDate := time.Now()
	y, m, d := endDate.Date()
	end = fmt.Sprintf("%d-%d-%d", y, m, d)

	// TODO: Make this configurable
	// Subtract 1 day
	startDate := endDate.AddDate(0, 0, -1)
	y, m, d = startDate.Date()
	start = fmt.Sprintf("%d-%d-%d", y, m, d)
	return
}

func getStartString(time time.Time) string {
	return fmt.Sprintf("-- %s %d/%d/%d --",
		time.Weekday(), time.Day(), time.Month(), time.Year())
}

// Print the given item in the regular format
func printDetailedTimeEntry(item toggl.DetailedTimeEntry) error {
	dur, err := getDuration(item.Duration)

	if err != nil {
		return err
	}

	startTime, endTime := item.Start.Local().Format("15:04"), item.End.Local().Format("15:04")
	startEnd := color.YellowString("(%s - %s)", startTime, endTime)
	fmt.Printf("%s %s - %s\n", startEnd, item.Description, dur.String())
	return nil
}

// Print the given item in the regular format
// This needs heeeaaavvvyyyyy refactoring to reduce duplication from printDetailedTimeEntry above
func printTimeEntry(item toggl.TimeEntry) error {
	if item.IsRunning() {
		startTime := item.Start.Local().Format("15:04")
		duration, err := getDuration(item.Duration)

		if err != nil {
			return err
		}

		fmt.Printf("%s %s - %d\n", startTime, item.Description, duration)
	}

	return nil
}

// Converts the duration in 'ms' to long-form
func getDuration(ms int64) (time.Duration, error) {
	// If the ms value is negative then it's actually the time since epoch
	if ms < 0 {
		second := ms + time.Now().Unix()
		return time.ParseDuration(fmt.Sprintf("%ds", second))
	}
	return time.ParseDuration(fmt.Sprintf("%dms", ms))
}
