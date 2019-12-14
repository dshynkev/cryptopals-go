package oracle

import (
	"crypto/aes"
	"crypto/rand"

	"cryptopals/s1/q7/ecb"
	"cryptopals/s2/q10/cbc"
	"cryptopals/s2/q11/detect"
)

const (
	BlockSize = aes.BlockSize

	MinPadLen = 5
	MaxPadLen = 10

	StartPadLenOffset = 0
	StartPadOffset    = StartPadLenOffset + 1
	EndPadLenOffset   = StartPadOffset + MaxPadLen
	EndPadOffset      = EndPadLenOffset + 1
	KeyOffset         = EndPadOffset + MaxPadLen
	IvOffset          = KeyOffset + BlockSize
	ModeOffset        = IvOffset + BlockSize
	EndOffset         = ModeOffset + 1

	RandomDataSize = EndOffset
)

// This does not result in a uniform distribution,
// but we don't truly care in this case.
func fit(in, min, max int) int {
	return in%(max-min) + min
}

type RandomEncryptor struct {
	LastBlockMode detect.AesBlockMode
	LastOutput    []byte
}

func (re *RandomEncryptor) Ask(in []byte) []byte {
	var randbuf = make([]byte, RandomDataSize)
	rand.Read(randbuf)

	var startPadLen = fit(int(randbuf[StartPadLenOffset]), MinPadLen, MaxPadLen)
	var endPadLen = fit(int(randbuf[EndPadLenOffset]), MinPadLen, MaxPadLen)

	var paddedLen = len(in) + startPadLen + endPadLen
	var padded = make([]byte, paddedLen)

	copy(
		padded[0:startPadLen],
		randbuf[StartPadOffset:(StartPadOffset+startPadLen)],
	)
	copy(
		padded[startPadLen:(paddedLen-endPadLen)],
		in,
	)
	copy(
		padded[(paddedLen-endPadLen):paddedLen],
		randbuf[EndPadOffset:(EndPadOffset+endPadLen)],
	)

	var key = randbuf[KeyOffset:(KeyOffset + BlockSize)]
	var iv = randbuf[IvOffset:(IvOffset + BlockSize)]

	if randbuf[ModeOffset]%2 == 0 {
		re.LastBlockMode = detect.EcbMode
		re.LastOutput, _ = ecb.Encrypt(padded, key)
	} else {
		re.LastBlockMode = detect.CbcMode
		re.LastOutput, _ = cbc.Encrypt(padded, key, iv)
	}

	return re.LastOutput
}

func NewRandomEncryptor() *RandomEncryptor {
	// zero-initialized is fine in this case
	return &RandomEncryptor{}
}
