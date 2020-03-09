package main

import (
   "encoding/asn1"
   "fmt"
   "os"
)

func main() {
   val := 13
   fmt.Println("Before marshal/unmarshal: ", val)
   mdata, err := asn1.Marshal(val)
   checkError(err)

   var n int
   _, err1 := asn1.Unmarshal(mdata, &n)
   checkError(err1)

   s := "hello"
   msdata, _ := asn1.Marshal(s)
   
   var newstr string
   asn1.Unmarshal(msdata, &newstr)

   fmt.Println("After marshal/unmarshal: ", n)
   fmt.Println("After marshal/unmarshal: ", newstr)
}

func checkError(err error) {
   if err != nil {
      fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
      os.Exit(1)
   }
}
