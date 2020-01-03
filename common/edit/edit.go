package edit

func Fill(in []byte, value byte) {
	for i := 0; i < len(in); i++ {
		in[i] = value
	}
}

func Expunge(in []byte, char byte) []byte {
	var i, N = 0, len(in)

	// Skip to the first occurence of char.
	for ; i < N; i++ {
		if in[i] == char {
			break
		}
	}
	if i == N {
		return in
	}

	var lag = 0

move:
	for {
		for in[i+lag] == char {
			lag += 1
			if i+lag >= N {
				break move
			}
		}
		in[i] = in[i+lag]
		i += 1
	}

	return in[:N-lag]
}
