package edit_test

import (
	"testing"

	"cryptopals/common/edit"
	"cryptopals/common/test"
)

func TestExpunge(t *testing.T) {
	in := []byte("a string\nwith\n\nnewlines\n")
	want := []byte("a stringwithnewlines")

	got := edit.Expunge(in, '\n')
	test.Test(t, want, got)
}
