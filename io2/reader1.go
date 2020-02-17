package main

import (
   "fmt"
   "os"
   "io"
)

type alphaReader string

func (a alphaReader) Read(p []byte) (int, error) {
   count := 0
   for i := 0; i < len(a); i++ {
      if (a[i] >= 'A' && a[i] <= 'Z') || (a[i] >= 'a' && a[i] <= 'z') {
         p[i] = a[i]
      }
      count++
   }
   return count, io.EOF
}

func main() {
   str := alphaReader("Hello! Where is the sun?")
   io.Copy(os.Stdout, &str)
   fmt.Println()
}
