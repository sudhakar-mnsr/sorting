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
