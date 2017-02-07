package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(assetFS()))
	// http.Handle("/", http.FileServer(http.Dir("/tmp/static/")))
	http.ListenAndServe(":8000", nil)
}
