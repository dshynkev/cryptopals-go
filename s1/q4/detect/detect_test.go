package detect

import (
	"bufio"
	"os"
	"testing"

	"cryptopals/common"
)

type result struct {
	idx   int
	bytes []byte
}

var want = result{
	idx:   170,
	bytes: []byte("Now that the party is jumping\n"),
}

func TestSBXor(t *testing.T) {
	var lines [][]byte

	f, _ := os.Open("lines.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, common.HexToRaw(scanner.Bytes()))
	}

	out, idx, _ := SBXor(lines, false)
	got := result{idx: idx, bytes: out}

	common.Test(t, want, got)
}
