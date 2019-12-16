package cbc_test

import (
	"io/ioutil"
	"testing"

	"cryptopals/common"
	"cryptopals/s2/q10/cbc"
)

var key = []byte("YELLOW SUBMARINE")
var iv = make([]byte, cbc.BlockSize)

var cipher, _ = ioutil.ReadFile("cipher.txt")
var plain, _ = ioutil.ReadFile("plain.txt")

func TestEncrypt(t *testing.T) {
	got, err := cbc.Encrypt(plain, key, iv)
	common.Test(t, nil, err)
	common.Test(t, cipher, got)
}

func TestDecrypt(t *testing.T) {
	got, err := cbc.Decrypt(cipher, key, iv)
	common.Test(t, nil, err)
	common.Test(t, plain, got)
}

func init() {
	cipher = common.Expunge(cipher, '\n')
	cipher = common.B64ToRaw(cipher)
}
