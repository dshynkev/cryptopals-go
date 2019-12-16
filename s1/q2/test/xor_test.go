package xor_test

import (
	"testing"

	"cryptopals/common"
	"cryptopals/s1/q2/xor"
)

var in1 = []byte("1c0111001f010100061a024b53535009181c")
var in2 = []byte("686974207468652062756c6c277320657965")
var want = []byte("746865206b696420646f6e277420706c6179")

func TestXor(t *testing.T) {
	b1 := common.HexToRaw(in1)
	b2 := common.HexToRaw(in2)

	got := xor.Xor(b1, b2)

	got = common.RawToHex(got)

	common.Test(t, want, got)
}
