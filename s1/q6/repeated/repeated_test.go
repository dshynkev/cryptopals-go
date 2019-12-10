package repeated

import (
	"bytes"
	"testing"

	"cryptopals/common"
)

var in = []byte(
	`0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a2622632427276527` +
		`2a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f`,
)
var want = []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")

func TestEditDistance(t *testing.T) {
	in1 := []byte("this is a test")
	in2 := []byte("wokka wokka!!!")
	want := 37

	got := EditDistance(in1, in2)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestBestKeylen(t *testing.T) {
	want := 3
	got := BestKeylen(common.HexToRaw(in), false)
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}

func TestBreak(t *testing.T) {
	got, _ := Break(common.HexToRaw(in), false)
	if !bytes.Equal(got, want) {
		t.Fatalf("\n= have =\n%s\n= want =\n%s", got, want)
	}
}
