package repeated

func Xor(in, key []byte) []byte {
	var K = len(key)

	for i := range in {
		in[i] ^= key[i%K]
	}

	return in
}
