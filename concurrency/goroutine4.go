package main

import "fmt"

func main() {
   start := 0
   stop := 300
   delta := 10
   go func() {
      for i := start; i <= stop; i += delta {
         fmt.Println(i)
      }
   }()

   fmt.Scanln()
}

