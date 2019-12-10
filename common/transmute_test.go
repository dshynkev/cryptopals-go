package common

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestExpunge(t *testing.T) {
	in := []byte("a string\nwith\n\nnewlines\n")
	want := []byte("a stringwithnewlines")

	got := Expunge(in, '\n')
	if !bytes.Equal(got, want) {
		t.Fatalf("\n= got =\n%s\n= want =\n%s\n", got, want)
	}
}
