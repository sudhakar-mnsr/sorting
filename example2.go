// Sample program to show how to use the http trace with a unique Client
// and Transport.
package main

import (
	"log"
	"net/http"
	"net/http/httptrace"
)

// transport is an http.RoundTripper that keeps track of the in-flight
// request and implements hooks to report HTTP tracing events.
type transport struct {
	current *http.Request
}

// RoundTrip wraps http.DefaultTransport.RoundTrip to keep track
// of the current request.
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	t.current = req
	return http.DefaultTransport.RoundTrip(req)
}

// GotConn prints whether the connection has been used previously
// for the current request.
func (t *transport) GotConn(info httptrace.GotConnInfo) {
	log.Printf("Connection reused for %v? %v\n", t.current.URL, info.Reused)
}

func main() {

	// Create a new request for the call.
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		log.Fatalln(err)
	}
	// Create the transport value we are binding event to.
	var t transport

	// Create a ClientTrace value for the events we care about.
	trace := httptrace.ClientTrace{
		GotConn: t.GotConn,
	}
