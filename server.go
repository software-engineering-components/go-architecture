package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {}

func main() {
    http.HandleFunc("/time", headers)
    http.ListenAndServe(":8795", nil)
}