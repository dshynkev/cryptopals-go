package ecb_test

import (
	"io/ioutil"
	"testing"

	"cryptopals/common"
	"cryptopals/s2/q12/ecb"
	"cryptopals/s2/q12/oracle"
)

func TestBlockMode(t *testing.T) {
	want, _ := ioutil.ReadFile("plain.txt")
	secret, _ := ioutil.ReadFile("cipher.txt")

	secret = common.Expunge(secret, '\n')
	secret = common.B64ToRaw(secret)

	enc := oracle.NewEncryptor(secret)
	got := ecb.Break(enc)

	common.Test(t, want, got)
}
