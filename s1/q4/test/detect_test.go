package detect_test

import (
	"bufio"
	"os"
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/test"
	"cryptopals/s1/q4/detect"
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
		lines = append(lines, convert.HexToRaw(scanner.Bytes()))
	}

	out, idx, _ := detect.SBXor(lines, false)
	got := result{idx: idx, bytes: out}

	test.Test(t, want, got)
}
