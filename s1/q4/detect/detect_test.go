package detect

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"cryptopals/common"
)

var wantOut = []byte("Now that the party is jumping\n")
var wantIdx = 170

func TestSBXor(t *testing.T) {
	var lines [][]byte

	f, _ := os.Open("lines.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, common.HexToRaw(scanner.Bytes()))
	}

	out, idx, _ := SBXor(lines, false)
	if idx != wantIdx || !bytes.Equal(out, wantOut) {
		t.Fatalf("got (%d, %s), want (%d, %s)", idx, out, wantIdx, wantOut)
	}
}
