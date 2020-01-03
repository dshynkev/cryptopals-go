package repeated

func Xor(in, key []byte) []byte {
	var out = make([]byte, len(in))

	var K = len(key)

	for i := range in {
		out[i] = in[i] ^ key[i%K]
	}

	return out
}
