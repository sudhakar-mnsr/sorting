package main

import (
   "fmt"
)

func main() {
ch := make(chan int, 10)
makeEvenNums(4, ch)

fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
fmt.Println(<-ch)
}
