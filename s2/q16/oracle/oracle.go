package oracle

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"

	"cryptopals/s2/q10/cbc"
)

func payloadFor(data []byte) []byte {
	var buf bytes.Buffer

	buf.WriteString("comment1=cooking%20MCs;userdata=")

	for i := 0; i < len(data); i++ {
		switch data[i] {
		case ';':
			buf.WriteString("%3B")
		case '=':
			buf.WriteString("%3D")
		default:
			buf.WriteByte(data[i])
		}
	}

	buf.WriteString(";comment2=%20like%20a%20pound%20of%20bacon")

	return buf.Bytes()
}

type KeyValueEncryptor struct {
	Key         []byte
	Iv          []byte
	Secret      []byte
	LastEncoded []byte
}

func (e *KeyValueEncryptor) Encrypt(data []byte) []byte {
	e.LastEncoded = payloadFor(data)
	out, err := cbc.Encrypt(e.LastEncoded, e.Key, e.Iv)
	if err != nil {
		panic(err)
	}

	return out
}

func (e *KeyValueEncryptor) Decrypt(ciphertext []byte) []byte {
	out, err := cbc.Decrypt(ciphertext, e.Key, e.Iv)
	if err != nil {
		panic(err)
	}

	if bytes.Contains(out, []byte(";admin=true;")) {
		return e.Secret
	} else {
		return nil
	}
}

func NewEncryptor() *KeyValueEncryptor {
	kve := new(KeyValueEncryptor)

	kve.Iv = make([]byte, aes.BlockSize)
	kve.Key = make([]byte, aes.BlockSize)
	kve.Secret = make([]byte, aes.BlockSize)

	rand.Read(kve.Key)
	rand.Read(kve.Iv)
	rand.Read(kve.Secret)

	return kve
}
