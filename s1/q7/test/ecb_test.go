package ecb_test

import (
	"io/ioutil"
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/edit"
	"cryptopals/common/test"
	"cryptopals/s1/q7/ecb"
)

var key = []byte("YELLOW SUBMARINE")
var plain, _ = ioutil.ReadFile("plain.txt")
var cipher, _ = ioutil.ReadFile("cipher.txt")

func TestEncrypt(t *testing.T) {
	got, _ := ecb.Encrypt(plain, key)

	test.Test(t, cipher, got)
}

func TestDecrypt(t *testing.T) {
	got, _ := ecb.Decrypt(cipher, key)

	test.Test(t, plain, got)
}

func init() {
	cipher = edit.Expunge(cipher, '\n')
	cipher = convert.B64ToRaw(cipher)
}
