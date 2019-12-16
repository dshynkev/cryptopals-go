package ecb

import (
	"crypto/aes"

	"cryptopals/common/pkcs7"
)

const BlockSize = aes.BlockSize

func Encrypt(in, key []byte) ([]byte, error) {
	in = pkcs7.Pad(in, BlockSize)
	out := make([]byte, len(in))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	offset, n := 0, len(in)/BlockSize
	for i := 0; i < n; i++ {
		cipher.Encrypt(
			out[offset:(offset+BlockSize)], in[offset:(offset+BlockSize)],
		)
		offset += BlockSize
	}

	return out, nil
}

func Decrypt(in, key []byte) ([]byte, error) {
	out := make([]byte, len(in))

	cipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	offset, n := 0, len(in)/BlockSize
	for i := 0; i < n; i++ {
		cipher.Decrypt(
			out[offset:(offset+BlockSize)], in[offset:(offset+BlockSize)],
		)
		offset += BlockSize
	}

	return pkcs7.Unpad(out)
}
