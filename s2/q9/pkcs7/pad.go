package pkcs7

const BlockSize = 16

func Pad(in []byte) []byte {
	padding := BlockSize - len(in)%BlockSize

	for i := 0; i < padding; i++ {
		in = append(in, byte(padding))
	}

	return in
}
