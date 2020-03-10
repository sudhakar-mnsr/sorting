package main

import (
        "encoding/json"
        "fmt"
        "net"
        "os"
)

type Person struct {
        Name  Name
        Email []Email
}

type Name struct {
        Family   string
        Personal string
}

type Email struct {
        Kind    string
        Address string
}

func (p Person) String() string {
   s := p.Name.Personal + " " + p.Name.Family
   for _, v := range p.Email {
      s += "\n" + v.Kind + ": " + v.Address
   }
   return s
}

func main() {
   service := "0.0.0.0:1200"
   serverPerson := Person{
      Name: Name{Family: "Server", Personal: "Jan"},
      Email: []Email{Email{Kind: "home", Address: "jan@server.name"},
                     Email{Kind: "work", Address: "j.newmarch@server.com"}}}
   tcpAddr, err := net.ResolveTCPAddr("tcp", service)
   checkError(err)

   listener, err := net.ListenTCP("tcp", tcpAddr)
   checkError(err)

   for {
      conn, err := listener.Accept()
      if err != nil {
         continue
      }
      encoder := json.NewEncoder(conn)
      decoder := json.NewDecoder(conn)

      for n := 0; n < 10; n++ {
         var clientPerson Person
         decoder.Decode(&clientPerson)
         fmt.Println(clientPerson.String())
         encoder.Encode(serverPerson)
      }
      conn.Close()
   }
} 

func checkError(err error) {
   if err != nil {
      fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}
