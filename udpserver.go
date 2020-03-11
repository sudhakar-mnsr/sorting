package main

import (
        "fmt"
        "net"
        "os"
        "time"
)

func main() {
   service := ":1200"
   udpAddr, err := net.ResolveUDPAddr("udp", service)
   checkError(err)

   conn, err := net.ListenUDP("udp", udpAddr)
   checkError(err)

   for {
      handleClient(conn)
   }
}
