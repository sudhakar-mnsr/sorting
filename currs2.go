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
   
   // loop to stay connected with client until client breaks
   for {
      // buffer for client command
      var cmdLine []byte
      // stream data using 4-byte chunks until io.EOF (\n)
      // The chunks are kept small to demonstrate streaming using io.Reader
      for {
         chunk := make([]byte, 4)
         n, err := conn.Read(chunk)
         if err != nil {
            // io.EOF may never happen since this is a stream
            if err == io.EOF {
               cmdLine, _ = appendBytes(cmdLine, chunk[:n]) // read remaining
               break
            }
            log.Println("connection read error:", err)
            return
         }
         if cmdLine, err = appendBytes(cmdLine, chunk[:n]); err == io.EOF {
            break
         }
      }
      cmd, param := parseCommand(string(cmdLine))
      if cmd == "" {
         if _, err := conn.Write([]byte("Invalid command\n")); err != nil {
            log.Println("failed to write:", err)
            return
         }
         continue
      }
      // Execute command
      switch strings.ToUpper(cmd) {
      case "GET":
         result := curr.Find(currencies, param)
         if len(result) == 0 {
            if _, err := conn.Write([]byte("Nothing found\n")); err != nil {
               log.Println("failed to write", err)
            }
            continue
         }
         // send each currency info as a line to the client with fmt.Fprintf()
         for _, cur := range result {
            _, err := conn.Write([]byte(fmt.Sprintf("%s %s %s %s\n",
                                 cur.Name, cur.Code, cur.Number, cur.Country,),
            )) 
            if err != nil {
               log.Println("failed to write response:", err)
               return
            }
         }
      default:
         if _, err := conn.Write([]byte("Invalid command\n")); err != nil {
            log.Println("failed to write:", err)
            return
         }
      }
   }
}  

func parseCommand(cmdLine string) (cmd, param string) {
	parts := strings.Split(cmdLine, " ")
	if len(parts) != 2 {
		return "", ""
	}
	cmd = strings.TrimSpace(parts[0])
	param = strings.TrimSpace(parts[1])
	return
}
