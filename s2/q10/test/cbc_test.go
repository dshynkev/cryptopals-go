package cbc_test

import (
	"io/ioutil"
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/edit"
	"cryptopals/common/test"
	"cryptopals/s2/q10/cbc"
)

var key = []byte("YELLOW SUBMARINE")
var iv = make([]byte, cbc.BlockSize)

var cipher, _ = ioutil.ReadFile("cipher.txt")
var plain, _ = ioutil.ReadFile("plain.txt")

func TestEncrypt(t *testing.T) {
	got, err := cbc.Encrypt(plain, key, iv)
	test.Test(t, nil, err)
	test.Test(t, cipher, got)
}

func TestDecrypt(t *testing.T) {
	got, err := cbc.Decrypt(cipher, key, iv)
	test.Test(t, nil, err)
	test.Test(t, plain, got)
}

func init() {
	cipher = edit.Expunge(cipher, '\n')
	cipher = convert.B64ToRaw(cipher)
}
