package oracle

type Encryptor interface {
	Encrypt([]byte) []byte
}

type EncryptDecryptor interface {
	Encrypt([]byte) []byte
	Decrypt([]byte) []byte
}
