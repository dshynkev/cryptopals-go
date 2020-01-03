package detect

import (
	sbxor "cryptopals/s1/q3/sbxor"
)

func SBXor(lines [][]byte, verbose bool) ([]byte, int, int) {
	var bestIdx int
	var bestKey byte
	var bestScore int

	for idx, line := range lines {
		key, score := sbxor.BestKey(line, verbose)
		if score > bestScore {
			bestIdx = idx
			bestKey = key
			bestScore = score
		}
	}

	var in = lines[bestIdx]
	var out = make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		out[i] = in[i] ^ bestKey
	}

	return out, bestIdx, bestScore
}
