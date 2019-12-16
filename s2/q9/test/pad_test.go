package pkcs7_test

import (
	"testing"

	"cryptopals/common"
	"cryptopals/common/pkcs7"
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
		common.Test(t, after[i], got)
	}
}

func TestUnpad(t *testing.T) {
	for i := 0; i < len(before); i++ {
		got, err := pkcs7.Unpad(after[i])
		common.Test(t, nil, err)
		common.Test(t, before[i], got)
	}
}

func TestUnpadCorrupted(t *testing.T) {
	var corrupted = []byte("BLUE SUBMARINE\x03\x03")
	_, err := pkcs7.Unpad(corrupted)
	common.Test(t, pkcs7.BadPadding, err)
}
