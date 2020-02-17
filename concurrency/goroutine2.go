package main

import "fmt"

func main() {
   go func(start, stop, delta int) {
      for i := start; i <= stop; i += delta {
         fmt.Println(i)
      }
   }(10,300,10)

   fmt.Scanln()
}

