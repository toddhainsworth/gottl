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

		for _, reportItem := range report.Data {
            time, err := getTime(reportItem.Duration)

            if err != nil {
                fmt.Println(err)
                return
            }

			fmt.Printf("%s - %v\n", reportItem.Description, time)
		}
	}
}

func getAPIKey() (string, error) {
	data, err := ioutil.ReadFile(os.ExpandEnv("$HOME/.gottl"))

	if err != nil {
		return "", err
	}

	str := string(data)
	return strings.TrimSpace(str), nil
}

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

func getTime(ms int64) (time.Duration, error) {
    return time.ParseDuration(fmt.Sprintf("%dms", ms))
}

func msToS(ms int64) int64 {
	return ms / 1000
}

func ppStruct(data interface{}) {
	fmt.Printf("%+v\n", data)
}
