# Gottl [![Build Status](https://travis-ci.org/toddhainsworth/gottl.svg?branch=master)](https://travis-ci.org/toddhainsworth/gottl)

Gets a list of time entries and their total duration for the last five days.
Uses [github.com/jason0x43/go-toggl](github.com/jason0x43/go-toggl) for a Toggl API interface

## Usage
1. Install [Go](https://golang.org/dl/)
1. Copy `.gottl.example` to `~/.gottl` and enter your details
* Your API key can be found within Toggl at the bottom of your profile settings
* Leave the workspace out for now, we can get that later
1. Install Gottl: `go install github.com/toddhainsworth/gottl`
1. Get your workspace IDs: `gottl --workspaces`
1. Add your desired workspace ID to the `~/.gottl` file from before
1. Get started `gottl`!

## Todo items
- [x] Basic functionality
- [x] Tests
- [ ] Custom date fetching
- [ ] Missing time notification
- [ ] Tests
- [x] Configurable account settings
- [ ] Multiple workspaces
- [ ] Tests
