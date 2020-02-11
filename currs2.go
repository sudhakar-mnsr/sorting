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
   
   // create a listener for provided network and host address
   ln, err := net.Listen(network, addr)
   if err != nil {
      log.Fatal("failed to create listener:", err)
   }
   defer ln.Close()
   log.Println("***** Global Currency Service *****")
   log.Printf("Service started: (%s) %s\n", network, addr)
   
   // connection loop handle incoming requests
   for {
      conn, err := ln.Accept()
      if err != nil {
         fmt.Println(err)
         if err := conn.Close(); err != nil {
            log.Println("failed to close listener:", err)
         }
         continue
      }
      log.Println("Connected to", conn.RemoteAddr())
      go handleConnection(conn)
   }
}

func handleConnection(conn net.Conn) {
defer func() {
   if err := conn.Close(); err != nil {
      log.Println("error closing connection:", err)
   }
}()
if _, err := conn.Write([]byte("Connected...\nUsage: GET <currency, country, or code>\n")); err != nil {
   log.Println("error writing:", err)
   return
}

// appendBytes is a func that simulates eof marker error
// since we will using streaming io on top of a streaming
// protocol, there may never be an actual eof marker. so
// this function simulates and io.EOF using \n

appendBytes := func(dest, src []byte) ([]byte, error) {
   for _, b := range src {
      if b == '\n' {
         return dest, io.EOF
      }
      dest = append(dest, b)
   }
   return dest, nil
}  
