package kv_test

import (
	"testing"

	"cryptopals/common"
	"cryptopals/s2/q16/kv"
	"cryptopals/s2/q16/oracle"
)

func TestBreak(t *testing.T) {
	enc := oracle.NewEncryptor()
	got := kv.Break(enc)
	common.Test(t, enc.Secret, got)
}
