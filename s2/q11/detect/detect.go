package detect

import (
	"bytes"
	"crypto/aes"

	"cryptopals/common/oracle"
)

type AesBlockMode int

const (
	EcbMode AesBlockMode = iota
	CbcMode
)

const BlockSize = aes.BlockSize

// Enough to always result in two identical consecutive blocks after padding:
// [ X padding ] [ BS-X testInput ] | [ BS testInput ] | [ BS testInput ] | ...
var testInput = bytes.Repeat([]byte{0x42}, 3*BlockSize)

func BlockMode(oracle oracle.Encryptor) AesBlockMode {
	out := oracle.Encrypt(testInput)

	if bytes.Equal(out[BlockSize:2*BlockSize], out[2*BlockSize:3*BlockSize]) {
		return EcbMode
	} else {
		return CbcMode
	}
}
