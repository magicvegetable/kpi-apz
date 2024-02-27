package main

import "fmt"
import "time"
import "net/http"
import "io"

type Time struct {
	time string
}

func (tm *Time) JSON() string {
	tm.time = time.Now().Format(time.RFC3339)
	return fmt.Sprintf("{\n\t\"time\": \"%v\"\n}", tm.time)
}

func handle(w http.ResponseWriter, r *http.Request) {
	var tm Time
	io.WriteString(w, tm.JSON())
}

func main() {
	http.HandleFunc("/time", handle)
	http.ListenAndServe(":8795", nil)
}
