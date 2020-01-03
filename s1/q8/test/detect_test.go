package detect_test

import (
	"bufio"
	"os"
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/test"
	"cryptopals/s1/q8/detect"
)

const want = 132

func TestDecrypt(t *testing.T) {
	f, _ := os.Open("in.txt")
	scanner := bufio.NewScanner(f)
	any := false

	for i := 0; scanner.Scan(); i++ {
		in := convert.HexToRaw(scanner.Bytes())
		detected := detect.IsEcb(in)
		if detected {
			test.Test(t, false, any)
			test.Test(t, want, i)
			any = true
		}
	}

	if !any {
		test.FailTest(t, want, -1)
	}
}
