package xor

func Xor(in1, in2 []byte) []byte {
	var out = make([]byte, len(in1))

	for i := 0; i < len(in1); i++ {
		out[i] = in1[i] ^ in2[i]
	}

	return out
}
