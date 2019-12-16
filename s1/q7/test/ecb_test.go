package ecb_test

import (
	"io/ioutil"
	"testing"

	"cryptopals/common"
	"cryptopals/s1/q7/ecb"
)

var key = []byte("YELLOW SUBMARINE")
var plain, _ = ioutil.ReadFile("plain.txt")
var cipher, _ = ioutil.ReadFile("cipher.txt")

func TestEncrypt(t *testing.T) {
	got, _ := ecb.Encrypt(plain, key)

	common.Test(t, cipher, got)
}

func TestDecrypt(t *testing.T) {
	got, _ := ecb.Decrypt(cipher, key)

	common.Test(t, plain, got)
}

func init() {
	cipher = common.Expunge(cipher, '\n')
	cipher = common.B64ToRaw(cipher)
}
