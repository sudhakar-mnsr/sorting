package main

import (
   "os"
   "io"
   "fmt"
)

type alphaWriter struct {
   src io.Writer
}

func newAlphaWriter(dest io.Writer) *alphaWriter {
    return &alphaWriter{dest}
}

func (a *alphaWriter) Write(p []byte) (int, error) {
   count := 0
   if len(p) == 0 {
      fmt.Println("Got here",len(p))
      return 0, nil
   }
   for i := 0; i < len(p); i++ {
      if (p[i] >= 'A' && p[i] <= 'Z') || (p[i] >= 'a' && p[i] <= 'z') {
         count++
         continue
      } else {
          p[i] = 0
      }
   }

   a.src.Write(p)
   return count, io.EOF
}

func main() {
   file, _ := os.Create("new.txt")
   alpha := newAlphaWriter(file)
   fmt.Fprint(alpha, "Hurray,,, You are already Streamed!")
}

