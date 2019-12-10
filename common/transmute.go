package common

func Expunge(in []byte, char byte) []byte {
	var i, N = 0, len(in)

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
	for ; ; i++ {
		for in[i+lag] == char {
			lag += 1
			if i+lag >= N {
				break move
			}
		}
		in[i] = in[i+lag]
	}

	return in[:N-lag]
}
