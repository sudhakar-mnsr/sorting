package main

import (
   "flag"
   "fmt"
   "io"
   "log"
   "net"
   "strings"
   curr "currency/lib"
)

var currencies = curr.Load("../data.csv")

func main() {
var addr string
var network string
flag.StringVar(&addr, "e", ":4040", "service endpoint [ip addr or sock path]")
flag.StringVar(&network, "n", "tcp", "network protocol [tcp or unix]")
flag.Parse()

// validate supported network protocols
switch network {
case "tcp","tcp4","tcp6","unix":
default:
   log.Fatalln("Unsupported network protocol:", network)
}
