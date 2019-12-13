package pkcs7

import (
	"testing"

	"cryptopals/common"
)

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
		got := Pad(before[i])
		common.Test(t, after[i], got)
	}
}

func TestUnpad(t *testing.T) {
	for i := 0; i < len(before); i++ {
		got := Unpad(after[i])
		common.Test(t, before[i], got)
	}
}
