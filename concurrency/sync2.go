package main

import (
   "sync"
   "time"
)

type Service struct {
   started bool
   stpCh chan struct{}
   mutex sync.Mutex
}

func (s *Service) Start() {
   s.stpCh = make(chan struct{})
   go func() {
      s.mutex.Lock()
      s.started = true
      s.mutex.Unlock()
      <-s.stpCh
   }()
}
