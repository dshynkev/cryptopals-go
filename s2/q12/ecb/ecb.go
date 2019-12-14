package ecb

import (
	"bytes"
	"errors"
	//"cryptopals/s2/q9/pkcs7"
)

type Oracle interface {
	Ask([]byte) []byte
}

const maxBlockSize = 64

var NotEcb = errors.New("not an AES-ECB oracle")

// getBlockSize returns:
// - the block size of an AES-ECB cipher used by oracle
// - whether the oracle, in fact, uses AES-ECB
func getBlockSize(oracle Oracle) (int, bool) {
	var input = bytes.Repeat([]byte{0x42}, maxBlockSize)

	for guess := 1; guess < maxBlockSize; guess *= 2 {
		out := oracle.Ask(input)
		if bytes.Equal(out[0:guess], out[guess:(2*guess)]) {
			return guess, true
		}
	}

	return 0, false
}

// getUnpaddedSize returns the size of oracle's secret before padding
func getUnpaddedSize(oracle Oracle, blockSize int) int {
	var input = bytes.Repeat([]byte{0x42}, blockSize)
	var startLen = len(oracle.Ask(nil))

	for guess := 1; guess < blockSize; guess++ {
		out := oracle.Ask(input[:guess])
		if len(out) > startLen {
			return startLen - guess
		}
	}

	return 0
}

func fill(buf []byte, count int, value byte) {
	for i := 0; i < count; i++ {
		buf[i] = value
	}
}

// Break takes an oracle which performs the following operation on its input:
//  AES-256-ECB(input || unknown-string, fixed-random-key)
// and decrypts unknown-string with repeated calls to the oracle.
func Break(oracle Oracle) ([]byte, error) {
	blockSize, ok := getBlockSize(oracle)
	if !ok {
		return nil, NotEcb
	}
	secretSize := getUnpaddedSize(oracle, blockSize)

	var plaintext = make([]byte, blockSize+secretSize)
	fill(plaintext, blockSize, '\x42')

	for offset := 0; offset < secretSize; offset++ {
		block, blockOffset := offset/blockSize, offset%blockSize

		reference := oracle.Ask(plaintext[blockOffset+1 : blockSize])
		for guess := 0; guess < 256; guess++ {
			plaintext[blockSize+offset] = byte(guess)
			this := oracle.Ask(plaintext[blockOffset+1 : blockSize+offset+1])
			matches := bytes.Equal(
				this[block*blockSize:(block+1)*blockSize],
				reference[block*blockSize:(block+1)*blockSize],
			)
			if matches {
				break
			}
		}
	}

	return plaintext[blockSize:], nil
}
