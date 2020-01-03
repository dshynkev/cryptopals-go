package repeated_test

import (
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/test"
	"cryptopals/s1/q6/repeated"
)

var in = []byte("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
var want = []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

func TestEditDistance(t *testing.T) {
	in1 := []byte("this is a test")
	in2 := []byte("wokka wokka!!!")
	want := 37

	got := repeated.EditDistance(in1, in2)
	test.Test(t, want, got)
}

func TestBestKeylen(t *testing.T) {
	want := 3
	got := repeated.BestKeylen(convert.HexToRaw(in), false)
	test.Test(t, want, got)
}

func TestBreak(t *testing.T) {
	got, _ := repeated.Break(convert.HexToRaw(in), false)
	test.Test(t, want, got)
}
