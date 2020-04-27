// Sample program to show how the httptrace package provides a number
// of hooks to gather information during an HTTP round trip about a
// variety of events.
package main

import (
	"log"
	"net/http"
	"net/http/httptrace"
)

func main() {

	// Create a new request for the call.
	req, err := http.NewRequest("GET", "http://goinggo.net", nil)
	if err != nil {
		log.Fatalln(err)
	}
