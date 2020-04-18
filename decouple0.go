package main

import (
   "errors"
   "fmt"
   "io"
   "math/rand"
   "time"
)

func init() {
   rand.Seed(time.Now().UnixNano())
}

// Data is the structure of the data we are copying
type Data struct {
   Line string
}

// Xenia is a system we need to pull data from.
type Xenia struct {
   Host string
   Timeout time.Duration
}

// PUll knows how to pull data out of Xenia.
func (*Xenia) Pull(d *Data) error {
   switch rand.Intn(10) {
   case 1,9:
      return io.EOF
   case 5:
      return errors.New("Error reading data from Xenia")
   default:
      d.Line = "Data"
      fmt.Println("In:", d.Line)
      return nil
   }
} 

// Pillar is a system we need to store data into.
type Pillar struct {
   Host string
   Timeout time.Duration
}

