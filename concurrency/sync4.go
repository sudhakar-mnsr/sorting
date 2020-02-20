package main

import (
   "fmt"
   "sync"
)

type Service struct {
   started bool
   stpCh chan struct{}
   mutex sync.RWMutex
   cache map[int]string
}

func (s *Service) Start() {
   s.stpCh = make(chan struct{})
   s.cache = make(map[int]string)
   go func() {
      s.mutex.Lock()
      s.started = true
      s.cache[1] = "Hello World"
      s.cache[2] = "Hello Universe"
      s.cache[3] = "Hello Galaxy!"
      s.mutex.Unlock()
      <-stpCh
   }()
}
