package oracle

import (
	"crypto/rand"

	"cryptopals/common/random"
	"cryptopals/s1/q7/ecb"
)

const BlockSize = ecb.BlockSize

type EcbFixedKeyEncryptor struct {
	Key    []byte
	Secret []byte
	Offset int
}

func (e *EcbFixedKeyEncryptor) Encrypt(in []byte) []byte {
	var combined = make([]byte, e.Offset+len(in)+len(e.Secret))

	rand.Read(combined[:e.Offset])
	copy(combined[e.Offset:], in)
	copy(combined[e.Offset+len(in):], e.Secret)

	var out, _ = ecb.Encrypt(combined, e.Key)
	return out
}

func NewEncryptor(secret []byte) *EcbFixedKeyEncryptor {
	var key = make([]byte, BlockSize)
	rand.Read(key)

	return &EcbFixedKeyEncryptor{
		Key:    key,
		Offset: int(random.Byte()),
		Secret: secret,
	}
}
