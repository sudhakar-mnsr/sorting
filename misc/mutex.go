// Sample program to show how to profile mutexes.
package mutex

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	// data is a slice that will be shared.
	data = make([]string, 1000)

	// rwMutex is used to define a critical section of code.
	rwMutex sync.RWMutex
)

// init is called prior to main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// TestMutexProfile creates goroutines that will content.
func TestMutexProfile(t *testing.T) {
	t.Log("Starting Test")

	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			writer()
			wg.Done()
		}()
