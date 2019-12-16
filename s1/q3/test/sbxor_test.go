package sbxor_test

import (
	"testing"

	"cryptopals/common"
	"cryptopals/s1/q3/sbxor"
)

var in1 = []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
var want = []byte("Cooking MC's like a pound of bacon")

func TestBreak(t *testing.T) {
	got, _ := sbxor.Break(common.HexToRaw(in1), false)

	common.Test(t, want, got)
}
