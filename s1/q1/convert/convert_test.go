package convert

import (
	"testing"

	"cryptopals/common"
)

var in = []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
var want = []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

func TestHexToB64(t *testing.T) {
	got := HexToB64(in)

	common.Test(t, want, got)
}
