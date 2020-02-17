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
   file1, _ := os.Open("reader1.go")
   alpha := newAlphaWriter(file)
   // fmt.Fprint(alpha, "Hurray,,, You are already Streamed!")
   io.Copy(alpha, file1) 
}

