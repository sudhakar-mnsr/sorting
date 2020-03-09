package main

import (
        "bytes"
        "encoding/asn1"
        "fmt"
        "io"
        "net"
        "os"
        "time"
)

func main() {
   if len(os.Args) != 2 {
      fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
      os.Exit(1)
   }
   service := os.Args[1]

   conn, err := net.Dial("tcp", service)
   checkError(err)

   result, err := readFully(conn)
   checkError(err)

   var newtime time.Time
   _, err1 := asn1.Unmarshal(result, &newtime)
   checkError(err1)

   fmt.Println("After marshal/unmarshal: ", newtime.String())

   os.Exit(0)
}
