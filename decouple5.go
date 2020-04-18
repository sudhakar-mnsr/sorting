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

type Puller interface {
   Pull(d *Data) error
}

type Storer interface {
   Store(d *Data) error
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

// Store knows how to store data into Pillar
func (*Pillar) Store(d *Data) error {
   fmt.Println("Out:", d.Line)
   return nil
}

// Pull knows how to pull bulks of data from Xenia
func pull(p Puller, data []Data) (int, error) {
   for i := range data {
      if err := p.Pull(&data[i]); err != nil {
         return i, err
      }
   }
   return len(data), nil
}

// store knows how to store bulks of data into Pillar
func store(s Storer, data []Data) (int, error) {
   for i := range data {
      if err := s.Store(&data[i]); err !=nil {
         return i, err
      }
   }
   return len(data), nil
}

func Copy(sys *System, batch int) error {
   data := make([]Data, batch)
   for {
      i, err := pull(sys, data)
      if i > 0 {
         if _, err := store(sys, data[:i]); err != nil {
            return err
         }
      }

      if err != nil {
         return err
      }
   }
}

func main() {
   sys := System{
             Puller: &Xenia{
                    Host: "localhost:8000",
                    Timeout: time.Second,
             },
             Storer: &Pillar{
                    Host: "localhost:9000",
                    Timeout: time.Second,
             },
   }
   if err := Copy(&sys, 3); err != io.EOF {
      fmt.Println(err)
   }
} 
