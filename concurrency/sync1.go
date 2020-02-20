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

