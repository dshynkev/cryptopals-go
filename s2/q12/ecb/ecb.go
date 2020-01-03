package ecb

import (
	"bytes"

	"cryptopals/common/edit"
	"cryptopals/common/oracle"
	"cryptopals/common/oracle/ecb"
)

// Break takes an oracle which performs the following operation on its input:
//  AES-256-ECB(input || unknown-string, fixed-random-key)
// and decrypts unknown-string with repeated calls to the oracle.
func Break(oracle oracle.Encryptor) []byte {
	layout := ecb.GetBlockLayout(oracle)
	secretSize := layout.BlockSize*layout.BlockCount - layout.Padding

	var plaintext = make([]byte, layout.BlockSize+secretSize)
	edit.Fill(plaintext[:layout.BlockSize], '\x42')

	for offset := 0; offset < secretSize; offset++ {
		block, blockOffset := offset/layout.BlockSize, offset%layout.BlockSize

		reference := oracle.Encrypt(plaintext[blockOffset+1 : layout.BlockSize])
		for guess := 0; guess < 256; guess++ {
			plaintext[layout.BlockSize+offset] = byte(guess)
			this := oracle.Encrypt(plaintext[blockOffset+1 : layout.BlockSize+offset+1])
			matches := bytes.Equal(
				this[block*layout.BlockSize:(block+1)*layout.BlockSize],
				reference[block*layout.BlockSize:(block+1)*layout.BlockSize],
			)
			if matches {
				break
			}
		}
	}

	return plaintext[layout.BlockSize:]
}
