package ctr

import (
	"crypto/aes"
	"errors"
)

const (
	BlockSize = 16

	NonceSize   = 7
	CounterSize = 9
)

var BadNonceSize = errors.New("bad nonce size")

func increment(counter []byte) {
	var last = len(counter) - 1
	counter[last] += 1
	for counter[last] == 0 && last >= 0 {
		last -= 1
		counter[last] += 1
	}
}

func xor(dst, src1, src2 []byte) {
	var length int
	if len(src1) < len(src2) {
		length = len(src1)
	} else {
		length = len(src2)
	}

	for i := 0; i < length; i++ {
		dst[i] = src1[i] ^ src2[i]
	}
}

func doCtr(in, key, nonce []byte) ([]byte, error) {
	if len(nonce) != NonceSize {
		return nil, BadNonceSize
	}

	var cipher, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	var counter = make([]byte, NonceSize+CounterSize)
	copy(counter[CounterSize:], nonce)

	var keystream = make([]byte, BlockSize)
	var out = make([]byte, len(in))

	for offset := 0; offset < len(in); offset += BlockSize {
		cipher.Encrypt(keystream, counter)
		xor(out[offset:], in[offset:], keystream)
		increment(counter[:CounterSize])
	}

	return out, nil
}

func Encrypt(in, key, nonce []byte) ([]byte, error) {
	return doCtr(in, key, nonce)
}

func Decrypt(in, key, nonce []byte) ([]byte, error) {
	return doCtr(in, key, nonce)
}
