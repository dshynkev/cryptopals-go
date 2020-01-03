package pad_test

import (
	"testing"

	"cryptopals/common/pkcs7"
	"cryptopals/common/test"
)

const blockSize = 16

var before = [][]byte{
	[]byte("BLUE SUBMARINE"),
	[]byte("YELLOW SUBMARINE"),
}

var after = [][]byte{
	[]byte("BLUE SUBMARINE\x02\x02"),
	[]byte("YELLOW SUBMARINE\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10"),
}

func TestPad(t *testing.T) {
	for i := 0; i < len(before); i++ {
		got := pkcs7.Pad(before[i], blockSize)
		test.Test(t, after[i], got)
	}
}
