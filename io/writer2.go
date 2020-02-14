package main

import "fmt"

type channelWriter struct {
   Channel chan byte
}

func NewChannelWriter() *channelWriter {
   return &channelWriter{
           Channel: make(chan byte, 1024),
   }
}
