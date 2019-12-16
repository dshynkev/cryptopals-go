package oracle

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"net/url"

	"cryptopals/s1/q7/ecb"
)

// We roll our own thing here because url.Values does not preserve order,
// and this only works if role comes last. We only escape the bare minimum.
func profileFor(email []byte) []byte {
	var buf bytes.Buffer

	buf.WriteString("email=")

	for i := 0; i < len(email); i++ {
		switch email[i] {
		case '%':
			buf.WriteString("%25")
		case '&':
			buf.WriteString("%26")
		default:
			buf.WriteByte(email[i])
		}
	}

	buf.WriteString("&uid=10&role=user")

	return buf.Bytes()
}

type KeyValueEncryptor struct {
	Key         []byte
	Secret      []byte
	LastEncoded []byte
}

func (e *KeyValueEncryptor) Encrypt(email []byte) []byte {
	e.LastEncoded = profileFor(email)
	out, _ := ecb.Encrypt(e.LastEncoded, e.Key)
	return out
}

func (e *KeyValueEncryptor) Decrypt(ciphertext []byte) []byte {
	out, _ := ecb.Decrypt(ciphertext, e.Key)
	parsed, _ := url.ParseQuery(string(out))

	if len(parsed["role"]) == 1 && parsed["role"][0] == "admin" {
		return e.Secret
	} else {
		return nil
	}
}

func NewEncryptor() *KeyValueEncryptor {
	kve := new(KeyValueEncryptor)

	kve.Key = make([]byte, aes.BlockSize)
	kve.Secret = make([]byte, aes.BlockSize)

	rand.Read(kve.Key)
	rand.Read(kve.Secret)

	return kve
}
