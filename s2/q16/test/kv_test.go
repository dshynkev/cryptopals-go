package kv_test

import (
	"testing"

	"cryptopals/common/test"
	"cryptopals/s2/q16/kv"
	"cryptopals/s2/q16/oracle"
)

func TestBreak(t *testing.T) {
	enc := oracle.NewEncryptor()
	got := kv.Break(enc)
	test.Test(t, enc.Secret, got)
}
