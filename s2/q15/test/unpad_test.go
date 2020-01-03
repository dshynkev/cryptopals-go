package unpad_test

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

func TestUnpad(t *testing.T) {
	for i := 0; i < len(before); i++ {
		got, err := pkcs7.Unpad(after[i], blockSize)
		test.Test(t, nil, err)
		test.Test(t, before[i], got)
	}
}

func TestUnpadCorrupted(t *testing.T) {
	var corrupted = [][]byte{
		[]byte{},
		[]byte("BLUE SUBMARINE"),
		[]byte("BLUE SUBMARINE\x03\x03"),
	}

	for i := 0; i < len(corrupted); i++ {
		_, err := pkcs7.Unpad(corrupted[i], blockSize)
		test.Test(t, pkcs7.BadPadding, err)
	}
}
