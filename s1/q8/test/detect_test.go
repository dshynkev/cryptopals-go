package detect_test

import (
	"bufio"
	"os"
	"testing"

	"cryptopals/common"
	"cryptopals/s1/q8/detect"
)

const want = 132

func TestDecrypt(t *testing.T) {
	f, _ := os.Open("in.txt")
	scanner := bufio.NewScanner(f)
	any := false

	for i := 0; scanner.Scan(); i++ {
		in := common.HexToRaw(scanner.Bytes())
		detected := detect.IsEcb(in)
		if detected {
			common.Test(t, false, any)
			common.Test(t, want, i)
			any = true
		}
	}

	if !any {
		common.FailTest(t, want, -1)
	}
}
