package ecb_test

import (
	"io/ioutil"
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/edit"
	"cryptopals/common/test"
	"cryptopals/s2/q12/ecb"
	"cryptopals/s2/q12/oracle"
)

func TestBlockMode(t *testing.T) {
	want, _ := ioutil.ReadFile("plain.txt")
	secret, _ := ioutil.ReadFile("cipher.txt")

	secret = edit.Expunge(secret, '\n')
	secret = convert.B64ToRaw(secret)

	enc := oracle.NewEncryptor(secret)
	got := ecb.Break(enc)

	test.Test(t, want, got)
}
