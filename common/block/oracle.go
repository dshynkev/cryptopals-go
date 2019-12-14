package block

import "bytes"

type EncryptOracle interface {
	Encrypt([]byte) []byte
}

type EncryptDecryptOracle interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}

// AES is defined only for {128,192,256} bits
const (
	MinBlockSize  = 16
	MaxBlockSize  = 32
	BlockSizeStep = 8
)

// static buffer twice the MaxBlockSize is enough for everything we want to do
var scratch = []byte{
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
}

// GetEcbBlockSize returns the block size of an AES-ECB cipher used by oracle
func GetEcbBlockSize(oracle EncryptOracle) int {
	var out = oracle.Encrypt(scratch)

	for guess := MinBlockSize; guess <= MaxBlockSize; guess += BlockSizeStep {
		if bytes.Equal(out[0:guess], out[guess:(2*guess)]) {
			return guess
		}
	}

	panic("not AES-ECB")
}

// GetEcbBlockCount returns the block count of oracle's response to empty request
func GetEcbBlockCount(oracle EncryptOracle, blockSize int) int {
	var out = oracle.Encrypt(nil)
	return len(out) / blockSize
}

// GetEcbBlockPadding returns the padding in oracle's response to empty request
func GetEcbBlockPadding(oracle EncryptOracle, blockSize int) int {
	var startLen = len(oracle.Encrypt(nil))

	for guess := 1; guess < blockSize; guess++ {
		if len(oracle.Encrypt(scratch[:guess])) > startLen {
			return guess
		}
	}

	return 0
}

// We often want to learn all the above at once
type EcbBlockLayout struct {
	Size    int
	Count   int
	Padding int
}

func GetEcbBlockLayout(oracle EncryptOracle) EcbBlockLayout {
	var ebl EcbBlockLayout

	ebl.Size = GetEcbBlockSize(oracle)
	ebl.Count = GetEcbBlockCount(oracle, ebl.Size)
	ebl.Padding = GetEcbBlockPadding(oracle, ebl.Size)

	return ebl
}
