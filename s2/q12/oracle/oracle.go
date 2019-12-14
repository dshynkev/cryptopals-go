package oracle

import (
	"crypto/aes"
	"crypto/rand"

	"cryptopals/s1/q7/ecb"
)

const BlockSize = aes.BlockSize

type EcbFixedKeyEncryptor struct {
	Key    []byte
	Secret []byte
}

func (e *EcbFixedKeyEncryptor) Encrypt(in []byte) []byte {
	var combined = make([]byte, len(in)+len(e.Secret))
	copy(combined, in)
	copy(combined[len(in):], e.Secret)
	var out, _ = ecb.Encrypt(combined, e.Key)
	return out
}

func NewEncryptor(secret []byte) *EcbFixedKeyEncryptor {
	var key = make([]byte, BlockSize)
	rand.Read(key)

	return &EcbFixedKeyEncryptor{
		Key:    key,
		Secret: secret,
	}
}
