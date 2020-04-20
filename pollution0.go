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

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
   // SPECIFIC IMPL
   return nil
}

// Stop shuts the server down
func (s *server) Stop() error {
   // SPECIFIC IMPL
   return nil
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
   // SPECIFIC IMPL
   return nil
}

func main() {
   // Create new server
   srv := NewServer("localhost")

   // Use the API
   srv.Start()
   srv.Stop()
   srv.Wait()
}

// Smells
// * The package declares an interface that matches the entire API of 
// * its own concrete type
// * The interface is exported but the concrete type is unexported.
// * The factory function returns the interface value with the unexported
// * concrete type value inside.
// * The interface can be removed and nothing changes for the user of API
// * The interface is not decoupleing the API from change.
