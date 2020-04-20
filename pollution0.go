// This is an example that creates interface pollution
// by improperly using an interface when one is not needed

package main

// Server defines a contract for tcp servers.
type Server interface {
   Start() error
   Stop() error
   Wait() error
}

// server is our server implementation
type server struct {
   host string
   // MORE FIELDS
}

// NewServer returns and interface value of type server
// with a server implementation
func NewServer(host string) Server {
   return &server{host}
}
