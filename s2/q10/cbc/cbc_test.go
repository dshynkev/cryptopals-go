package cbc

import (
	"io/ioutil"
	"testing"

	"cryptopals/common"
)

func TestDecrypt(t *testing.T) {
	var key = []byte("YELLOW SUBMARINE")
	var iv = make([]byte, BlockSize)

	in, _ := ioutil.ReadFile("in.txt")
	want, _ := ioutil.ReadFile("out.txt")

	in = common.Expunge(in, '\n')

	got, err := Decrypt(common.B64ToRaw(in), key, iv)
	common.Test(t, nil, err)
	common.Test(t, want, got)
}
