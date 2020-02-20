package main

import "time"

type service struct {
   started bool
   stpCh chan struct{}
}

func (s *Service) Start{} {
   s.stpCh = make (chan struct {})
   go func() {
      s.started = true
      <-s.stpCh
   }()
}


func (s *Service) Stop() {
   if s.started {
      s.started = false
      close(s.stpCh)
   }
}

func main() {
   s := &Service{}
   s.Start()
   time.Sleep(time.Second)
   s.Stop()
}
