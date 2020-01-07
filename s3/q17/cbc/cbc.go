package cbc

import (
	"crypto/aes"
	"cryptopals/common/pkcs7"
)

type CbcPaddingOracle interface {
	Encrypt() []byte
	CheckPadding([]byte) bool
}

const BlockSize = aes.BlockSize

func Break(oracle CbcPaddingOracle) []byte {
	// This is the read-only "original" ciphertext: c[i]
	ciphertext := oracle.Encrypt()

	// This is the "modified" ciphertext: s[i]
	scrap := make([]byte, len(ciphertext))
	copy(scrap, ciphertext)

	blockNum := len(ciphertext) / BlockSize
	// This is the recovered plaintext: p[i]
	//
	// We recover n-th block of plaintext from oracle's responses
	// to modifications of the *(n-1)-th* block of ciphertext.
	// As a corollary, the 0-th block cannot be recovered,
	// but, luckily, in AES-CBC it is just the IV.
	// In any case, the plaintext is a block shorter.
	plaintext := make([]byte, (blockNum-1)*BlockSize)

	for block := blockNum - 1; block > 0; block-- {
		// Restore the next block: it was scrambled in the last iteration
		copy(
			scrap[block*BlockSize:(block+1)*BlockSize],
			ciphertext[block*BlockSize:(block+1)*BlockSize],
		)
		// The core idea is this: we have
		//  AES-ENC(p[i] XOR c[i]) == c[i+B]
		// so after copying c to s we have
		//  AES-DEC(c[i+BS]) XOR s[i] == p[i]
		// But if we modify s so that s[i] != c[i], then
		//  AES-DEC(c[i+BS]) XOR s[i] == p[i] XOR (s[i] XOR c[i])
		// and the RHS here can be forced to any value with an appropriate s[i].
		// If that value is known, p[i] can be recovered as value XOR s[i] XOR c[i].
		for offset := 1; offset <= BlockSize; offset++ {
			// The fact that plaintext is a block shorter allows us to have
			// pleasantly uniform indices: p[idx] is recovered by iteration on s[idx].
			idx := block*BlockSize - offset
			// We want to detect when s[idx+BS] decrypts to offset.
			// Since we know p[idx+1...], we make the block tail (s[i] | i > idx)
			// such that each s[i+BS] decrypts to offset,
			// then padding is valid iff so does s[idx+BS].
			// This is true when (p[i] XOR c[i]) XOR s[i] = offset,
			// so s[i] = offset XOR p[i] XOR c[i] is the desired value.
			for i := idx + 1; i%BlockSize != 0; i++ {
				scrap[i] = byte(offset) ^ plaintext[i] ^ ciphertext[i]
			}
			for guess := 0; guess < 256; guess++ {
				scrap[idx] = byte(guess)
				// We only pass in s up to the block containing idx+BS
				// so that this block is considered last and checked for padding.
				if oracle.CheckPadding(scrap[:(block+1)*BlockSize]) {
					// If this is the last byte in this block, we may have got unlucky:
					// if the second-last byte is 0x02, then we have [... 0x02 ___]
					// and the padding is correct for both 0x01 and 0x02 as the last byte,
					// and similarly for the rest of 0x03...0x10 (albeit less likely).
					// However, we only want 0x01, so we try changing the next byte
					// to ensure that the padding still works: this is only true of 0x01.
					if offset == 1 {
						scrap[idx-1] += 1
						if !oracle.CheckPadding(scrap[:(block+1)*BlockSize]) {
							// We must restore the byte as further tries on idx will be made.
							// Note that this is unnecessary if the padding does check out:
							// in that case, this is the last iteration of the loop and idx-1
							// will be overwritten on the next one without being read.
							scrap[idx-1] -= 1
							continue
						}
					}
					// as before, p[i] = value XOR s[i] XOR c[i]
					plaintext[idx] = byte(offset^guess) ^ ciphertext[idx]
					break
				}
			}
		}
	}

	secret, _ := pkcs7.Unpad(plaintext, BlockSize)
	return secret
}
