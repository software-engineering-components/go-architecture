package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type JsonTime struct {
	Time string `json:"time"`
}

func getCurrentTime(w http.ResponseWriter, r *http.Request) {
	rfc3339 := time.Now().Format("2006-01-02 15:04:05.000")

	out, _ := json.Marshal(JsonTime{Time: rfc3339})

	fmt.Fprintf(w, string(out))
}

func main() {
	http.HandleFunc("/time", getCurrentTime)
	http.ListenAndServe(":8795", nil)
}
