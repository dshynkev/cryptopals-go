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

func Unpad(in []byte, blockSize int) ([]byte, error) {
	N := len(in)
	// The padded length must be a multiple of block size, but never 0:
	// PKCS7 requires a full block of padding in case input is block-aligned.
	if N == 0 || N%blockSize != 0 {
		return nil, BadPadding
	}

	// The last byte is certainly padding.
	// Each padding byte equals the padding length.
	padLength := in[N-1]
	// Padding defined to be between 1 and blockSize bytes.
	if padLength == 0 || int(padLength) > blockSize {
		return nil, BadPadding
	}

	inputLength := N - int(padLength)
	for i := inputLength; i < N; i++ {
		if in[i] != padLength {
			return nil, BadPadding
		}
	}

	return in[:inputLength], nil
}
