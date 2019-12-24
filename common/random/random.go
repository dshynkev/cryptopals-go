package random

import "crypto/rand"

const length = 16

var scratch [length]byte
var offset int = length

func ensureOffset() {
	if offset == length {
		rand.Read(scratch[:])
		offset = 0
	}
}

func Byte() byte {
	ensureOffset()
	offset += 1
	return scratch[offset-1]
}
