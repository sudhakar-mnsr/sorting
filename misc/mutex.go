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
