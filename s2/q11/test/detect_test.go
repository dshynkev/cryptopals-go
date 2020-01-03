package detect_test

import (
	"testing"

	"cryptopals/common/test"
	"cryptopals/s2/q11/detect"
	"cryptopals/s2/q11/oracle"
)

const N = 10

func TestBlockMode(t *testing.T) {
	re := oracle.NewEncryptor()

	for i := 0; i < N; i++ {
		mode := detect.BlockMode(re)
		test.Test(t, re.LastBlockMode, mode)
	}
}
