package pkcs7

const BlockSize = 16

func Pad(in []byte) []byte {
	padding := BlockSize - len(in)%BlockSize

	for i := 0; i < padding; i++ {
		in = append(in, byte(padding))
	}

	return in
}

func Unpad(in []byte) []byte {
	N := len(in)
	padding := int(in[N-1])
	return in[:N-padding]
}
