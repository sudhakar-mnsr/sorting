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
	// Create a ClientTrace value for the events we care about.
	trace := httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			log.Printf("Get Conn: %s\n", hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			log.Printf("Got Conn: %+v\n", connInfo)
		},

		DNSStart: func(dnsInfo httptrace.DNSStartInfo) {
			log.Printf("DNS Start Info: %+v\n", dnsInfo)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			log.Printf("DNS Done Info: %+v\n", dnsInfo)
		},

		ConnectStart: func(network, addr string) {
			log.Printf("Connect Start: %s, %s\n", network, addr)
		},
		ConnectDone: func(network, addr string, err error) {
			log.Printf("Connect Done: %s, %s, %v\n", network, addr, err)
		},
