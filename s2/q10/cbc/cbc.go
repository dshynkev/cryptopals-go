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
	if len(in) == 0 {
		return nil, nil
	}

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

func Encrypt(in, key, iv []byte) ([]byte, error) {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	in = pkcs7.Pad(in)

	N := len(in) / BlockSize
	out := make([]byte, len(in))
	tmp := make([]byte, BlockSize)

	copy(tmp, in[0:BlockSize])
	xor(tmp, iv)
	cipher.Encrypt(out[0:BlockSize], tmp)

	offset := BlockSize
	for i := 1; i < N; i++ {
		copy(tmp, in[offset:(offset+BlockSize)])
		xor(tmp, out[(offset-BlockSize):offset])
		cipher.Encrypt(out[offset:(offset+BlockSize)], tmp)

		offset += BlockSize
	}

	return out, nil
}
