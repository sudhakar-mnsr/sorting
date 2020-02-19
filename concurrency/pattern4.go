package main

import (
	"fmt"
	"strings"
)

func main() {
   data := []string{
      "The yellow fish swims slowly in the water",
      "The brown dog barks loudly after a drink from its water bowl",
      "The dark bird of prey lands on a small tree after hunting for fish",
   }
   
   histogram := make(map[string]int)
   
   words := words(data)
   for word := range words {
      histogram[word]++
   }
   
   for k, v := range histogram {
      fmt.Println("%s\t(%d)\n", k, v)
   }
}

// generator function that produces data
func words(data []string) <-chan string {
   out := make(chan string)
   go func() {
      defer close(out)
      for _, line := range data {
         words := strings.Split(line, " ")
         for _, word := range words {
            word = strings.ToLower(word)
            out <- word
         }
      }
   }()
   return out
}

