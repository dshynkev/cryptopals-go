package kv

import (
	"crypto/aes"
	"cryptopals/common/block"
)

const payload = `BLOCK WITH ERRORyoink:admin<true`

func Break(oracle block.EncryptDecryptOracle) []byte {
	var response = oracle.Encrypt([]byte(payload))

	response[2*aes.BlockSize+5] |= 1
	response[2*aes.BlockSize+11] |= 1

	return oracle.Decrypt(response)
}
