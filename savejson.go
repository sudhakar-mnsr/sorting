package main

import (
   "encoding/json"
   "fmt"
   "os"
)

type Person struct {
   Name Name
   Email []Email
}

type Name struct {
   Family string
   Personal string
}

type Email struct {
   Kind string
   Address string
}

func main() {
   person := Person{
      Name: Name{Family: "Manukonda", Personal: "Jan"},
      Email: []Email{Email{Kind: "home", Address: "jan@manukonda.name"},
      Email{Kind: "Work", Address: "jan.manukonda@boxhill.edu.au"}}}
   saveJSON(person.json", person)
}
