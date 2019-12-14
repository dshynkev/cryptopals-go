package kv_test

import (
	"testing"

	"cryptopals/common"
	"cryptopals/s2/q13/kv"
	"cryptopals/s2/q13/oracle"
)

func TestBreak(t *testing.T) {
	enc := oracle.NewEncryptor()
	got := kv.Break(enc)
	common.Test(t, enc.Secret, got)
}
