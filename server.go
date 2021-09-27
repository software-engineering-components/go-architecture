package main

import (
  "fmt"
  "net/http"
)

func getServerTime(w http.ResponseWriter, req *http.Request) {}

func main() {
  http.HandleFunc("/time", getServerTime)
  http.ListenAndServe(":8795", nil)
}
