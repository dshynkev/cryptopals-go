package repeated_test

import (
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/test"
	"cryptopals/s1/q5/repeated"
)

var in1 = []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
var key = []byte("ICE")
var want = []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")

func TestXor(t *testing.T) {
	got := repeated.Xor(in1, key)
	test.Test(t, want, convert.RawToHex(got))
}
