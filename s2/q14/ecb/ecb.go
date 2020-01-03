package ecb

import (
	"bytes"
	"log"

	"cryptopals/common/edit"
	"cryptopals/common/oracle"
	"cryptopals/common/oracle/ecb"
)

// getInsertOffset finds the offset of user input in the final oracle payload.
func getInsertOffset(oracle oracle.Encryptor, insertBlock, blockSize int) int {
	var scratch = make([]byte, 3*blockSize)
	for i := 0; i < blockSize; i++ {
		scratch[i] = 0x43
		out := oracle.Encrypt(scratch)
		if !bytes.Equal(
			out[insertBlock*blockSize:(insertBlock+1)*blockSize],
			out[(insertBlock+1)*blockSize:(insertBlock+2)*blockSize],
		) {
			return insertBlock*blockSize - i
		}
		scratch[i] = 0
	}

	return -1
}

// Break takes an oracle which performs the following operation on its input:
//  AES-256-ECB(random-padding || input || unknown-string, fixed-random-key)
// and decrypts unknown-string with repeated calls to the oracle.
func Break(oracle oracle.Encryptor) []byte {
	layout := ecb.GetBlockLayout(oracle)
	insertBlock := ecb.GetInsertionBlock(oracle, layout.BlockSize)
	insertOffset := getInsertOffset(oracle, insertBlock, layout.BlockSize)
	log.Printf("%+v %d %d", layout, insertBlock, insertOffset)

	secretSize := layout.BlockSize*layout.BlockCount - layout.Padding - insertOffset

	startPadding := 2*layout.BlockSize - insertOffset%layout.BlockSize
	var plaintext = make([]byte, startPadding+secretSize)
	edit.Fill(plaintext[:startPadding], '\x42')

	for offset := 0; offset < secretSize; offset++ {
		blockNum, blockOffset := offset/layout.BlockSize, offset%layout.BlockSize

		reference := oracle.Encrypt(plaintext[blockOffset+1 : startPadding])
		for guess := 0; guess < 256; guess++ {
			plaintext[startPadding+offset] = byte(guess)
			this := oracle.Encrypt(plaintext[blockOffset+1 : startPadding+offset+1])
			start := (insertBlock + blockNum) * layout.BlockSize
			matches := bytes.Equal(
				this[start:start+layout.BlockSize],
				reference[start:start+layout.BlockSize],
			)
			if matches {
				break
			}
		}
	}

	return plaintext[startPadding:]
}
