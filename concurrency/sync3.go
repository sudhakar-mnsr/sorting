package main

import (
   "sync"
   "time"
)

type Service struct {
   started bool
   stpCh chan struct{}
   sync.Mutex
}

func (s *Service) Start() {
   s.stpCh = make(chan struct{})
   go func() {
      s.Lock()
      s.started = true
      s.Unlock()
      <-s.stpCh
   }()
}
 
