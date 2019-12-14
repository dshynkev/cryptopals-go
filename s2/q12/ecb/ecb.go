package ecb

import (
	"bytes"

	"cryptopals/common/block"
)

func fill(buf []byte, count int, value byte) {
	for i := 0; i < count; i++ {
		buf[i] = value
	}
}

// Break takes an oracle which performs the following operation on its input:
//  AES-256-ECB(input || unknown-string, fixed-random-key)
// and decrypts unknown-string with repeated calls to the oracle.
func Break(oracle block.EncryptOracle) []byte {
	layout := block.GetEcbBlockLayout(oracle)
	secretSize := layout.Size*layout.Count - layout.Padding

	var plaintext = make([]byte, layout.Size+secretSize)
	fill(plaintext, layout.Size, '\x42')

	for offset := 0; offset < secretSize; offset++ {
		block, blockOffset := offset/layout.Size, offset%layout.Size

		reference := oracle.Encrypt(plaintext[blockOffset+1 : layout.Size])
		for guess := 0; guess < 256; guess++ {
			plaintext[layout.Size+offset] = byte(guess)
			this := oracle.Encrypt(plaintext[blockOffset+1 : layout.Size+offset+1])
			matches := bytes.Equal(
				this[block*layout.Size:(block+1)*layout.Size],
				reference[block*layout.Size:(block+1)*layout.Size],
			)
			if matches {
				break
			}
		}
	}

	return plaintext[layout.Size:]
}
