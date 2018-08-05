package main

import "testing"

const APIKEY = "ABC123"

func TestNewApp(t *testing.T) {
	app := NewApp(APIKEY)

	if key := app.APIKey; key != APIKEY {
		t.Fatalf("APIKeys do not match, expected \"%s\", got \"%s\"", APIKEY, key)
	}
}

func TestStartSession(t *testing.T) {
	app := NewApp(APIKEY)

	err := app.StartSession()
	if err != nil {
		t.Fatalf("Expected no error, instead got \"%v\"", err.Error())
	}

	app = App{}
	err = app.StartSession()
	if err.Error() != "API key is not present" {
		t.Fatalf("Expected error: \"API key is not present\", got: \"%v\"", err.Error())
	}
}

func TestPrintReport(t *testing.T) {
	// TODO(todd): test stdout
}
