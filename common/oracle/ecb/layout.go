package ecb

import (
	"bytes"

	"cryptopals/common/oracle"
)

// AES is defined only for {128,192,256} bits
const (
	MinBlockSize  = 16
	MaxBlockSize  = 32
	BlockSizeStep = 8
)

// this 128-byte static buffer is the most need for any operation
var scratch = []byte{
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
	0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42, 0x42,
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
func GetBlockSize(oracle oracle.Encryptor) int {
	var out = oracle.Encrypt(scratch[:])

	for guess := MinBlockSize; guess <= MaxBlockSize; guess += BlockSizeStep {
		for i := 1; (i+1)*guess <= len(out); i++ {
			if bytes.Equal(out[(i-1)*guess:i*guess], out[i*guess:(i+1)*guess]) {
				return guess
			}
		}
	}

	panic("not AES-ECB")
}

// GetBlockCount returns the block count of oracle's response to empty request
func GetBlockCount(oracle oracle.Encryptor, bsize int) int {
	var out = oracle.Encrypt(nil)
	return len(out) / bsize
}

// GetPadding returns the padding in oracle's response to empty request
func GetPadding(oracle oracle.Encryptor, bsize int) int {
	var startLen = len(oracle.Encrypt(nil))

	for guess := 1; guess < bsize; guess++ {
		if len(oracle.Encrypt(scratch[:guess])) > startLen {
			return guess
		}
	}

	return 0
}

// We often want to learn all the above at once
type BlockLayout struct {
	BlockSize  int
	BlockCount int
	Padding    int
}

func GetBlockLayout(oracle oracle.Encryptor) BlockLayout {
	var bl BlockLayout

	bl.BlockSize = GetBlockSize(oracle)
	bl.BlockCount = GetBlockCount(oracle, bl.BlockSize)
	bl.Padding = GetPadding(oracle, bl.BlockSize)

	return bl
}

// GetInsertionBlock returns the block number where oracle input is injected
func GetInsertionBlock(oracle oracle.Encryptor, bsize int) int {
	out := oracle.Encrypt(scratch)

	n := len(out) / bsize
	for i := 1; i+1 < n; i++ {
		if bytes.Equal(out[(i-1)*bsize:i*bsize], out[i*bsize:(i+1)*bsize]) {
			return i - 1
		}
	}

	panic("not AES-ECB")
}
