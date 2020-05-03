// Package bce shows a sample function that does not take into consideration
// the extra bounds checks that the compiler places in code for integrity. By
// using information provided by the ssa backend of the compiler, we can
// change the code to remove the checks and improve performance.
package bce

import "encoding/binary"

// go build -gcflags -d=ssa/check_bce/debug=1

func hash64(buffer []byte, seed uint64) uint64 {
	const (
		k0 = 0xD6D018F5
		k1 = 0xA2AA033B
		k2 = 0x62992FC1
		k3 = 0x30BC5B29
	)

	ptr := buffer

	hash := (seed + k2) * k0

	if len(ptr) >= 32 {
		v := [4]uint64{hash, hash, hash, hash}
