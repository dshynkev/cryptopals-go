package ecb

import (
	"crypto/aes"

	"cryptopals/s2/q9/pkcs7"
)

func Encrypt(in, key []byte) ([]byte, error) {
	in = pkcs7.Pad(in)
	out := make([]byte, len(in))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	offset, n := 0, len(in)/aes.BlockSize
	for i := 0; i < n; i++ {
		cipher.Encrypt(
			out[offset:(offset+aes.BlockSize)], in[offset:(offset+aes.BlockSize)],
		)
		offset += aes.BlockSize
	}

	return out, nil
}

func Decrypt(in, key []byte) ([]byte, error) {
	out := make([]byte, len(in))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	offset, n := 0, len(in)/aes.BlockSize
	for i := 0; i < n; i++ {
		cipher.Decrypt(
			out[offset:(offset+aes.BlockSize)], in[offset:(offset+aes.BlockSize)],
		)
		offset += aes.BlockSize
	}

	return pkcs7.Unpad(out), nil
}
