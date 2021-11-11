package main

import (
	"fmt"
	"net/http"
	"os"
)

var Hit int = 0

func main() {
	http.HandleFunc("/", Serve)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "80"
	}

	http.ListenAndServe(":"+httpPort, nil)
}

func Serve(w http.ResponseWriter, r *http.Request) {
	Hit++
	fmt.Fprintf(w, "User address:  %s\nHits: %d", r.RemoteAddr, Hit)
}
