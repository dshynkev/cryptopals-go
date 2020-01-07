package oracle

import (
	"crypto/rand"

	"cryptopals/common/convert"
	"cryptopals/common/pkcs7"
	"cryptopals/common/random"
	"cryptopals/s2/q10/cbc"
)

const BlockSize = cbc.BlockSize

type CbcPaddingOracle struct {
	Key    []byte
	Secret []byte
}

func (o *CbcPaddingOracle) Encrypt() []byte {
	var iv = make([]byte, BlockSize)
	rand.Read(iv)

	var encrypted, _ = cbc.Encrypt(o.Secret, o.Key, iv)
	return append(iv, encrypted...)
}

func (o *CbcPaddingOracle) CheckPadding(in []byte) bool {
	if len(in) < BlockSize {
		return false
	}
	var _, err = cbc.Decrypt(in[BlockSize:], o.Key, in[:BlockSize])
	return err != pkcs7.BadPadding
}

var _secrets = []string{
	"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
	"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
	"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
	"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
	"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
	"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
	"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
	"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
	"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
	"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
}

func NewEncryptor() *CbcPaddingOracle {
	var key = make([]byte, BlockSize)
	rand.Read(key)

	var secretIdx = int(random.Byte()) % len(_secrets)
	var secret = convert.B64ToRaw([]byte(_secrets[secretIdx]))

	return &CbcPaddingOracle{
		Key:    key,
		Secret: secret,
	}
}
