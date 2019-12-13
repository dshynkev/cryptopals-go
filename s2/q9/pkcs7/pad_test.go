package pkcs7

import (
	"testing"

	"cryptopals/common"
)

func TestPad(t *testing.T) {
	in := [][]byte{
		[]byte("BLUE SUBMARINE"),
		[]byte("YELLOW SUBMARINE"),
	}
	want := [][]byte{
		[]byte("BLUE SUBMARINE\x02\x02"),
		[]byte("YELLOW SUBMARINE\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10"),
	}

	for i := 0; i < len(in); i++ {
		got := Pad(in[i])
		common.Test(t, want[i], got)
	}
}
