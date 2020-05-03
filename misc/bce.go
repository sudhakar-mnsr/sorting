// Package bce shows a sample function that does not take into consideration
// the extra bounds checks that the compiler places in code for integrity. By
// using information provided by the ssa backend of the compiler, we can
// change the code to remove the checks and improve performance.
package bce

import "encoding/binary"

// go build -gcflags -d=ssa/check_bce/debug=1
