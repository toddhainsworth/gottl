package main

import (
  "fmt"
  "github.com/jason0x43/go-toggl"
)

const APIKEY = "<goes-here>"

func main() {
  toggl.DisableLog()
  session := toggl.OpenSession(APIKEY)
  account, err := session.GetAccount()

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(account.Data.BeginningOfWeek)
}
