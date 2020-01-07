package cbc_test

import (
	"testing"

	"cryptopals/common/test"
	"cryptopals/s3/q17/cbc"
	"cryptopals/s3/q17/oracle"
)

func TestBlockMode(t *testing.T) {
	enc := oracle.NewEncryptor()
	got := cbc.Break(enc)

	test.Test(t, enc.Secret, got)
}
