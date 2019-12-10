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

	out := lines[bestIdx]
	for i := 0; i < len(out); i++ {
		out[i] ^= bestKey
	}

	return out, bestIdx, bestScore
}
