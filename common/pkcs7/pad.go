package pkcs7

import "errors"

var BadPadding = errors.New("bad padding")

func Pad(in []byte, blockSize int) []byte {
	padding := blockSize - len(in)%blockSize

	for i := 0; i < padding; i++ {
		in = append(in, byte(padding))
	}

	return in
}

func Unpad(in []byte) ([]byte, error) {
	N := len(in)
	padding := in[N-1]

	end := N - int(padding)
	for i := end; i < N; i++ {
		if in[i] != padding {
			return nil, BadPadding
		}
	}

	return in[:end], nil
}
