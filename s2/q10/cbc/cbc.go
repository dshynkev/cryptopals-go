package cbc

import (
	"crypto/aes"

	"cryptopals/s2/q9/pkcs7"
)

const BlockSize = aes.BlockSize

func xor(out []byte, mask []byte) {
	for i := 0; i < len(mask); i++ {
		out[i] ^= mask[i]
	}
}

func Decrypt(in, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(in))
	N := len(in) / BlockSize

	cipher.Decrypt(out[0:BlockSize], in[0:BlockSize])
	xor(out[0:BlockSize], iv)

	offset := BlockSize
	for i := 1; i < N; i++ {
		cipher.Decrypt(
			out[offset:(offset+BlockSize)],
			in[offset:(offset+BlockSize)],
		)
		xor(
			out[offset:(offset+BlockSize)],
			in[(offset-BlockSize):offset],
		)

		offset += BlockSize
	}

	return pkcs7.Unpad(out), nil
}
