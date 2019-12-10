package sbxor

import "fmt"

func isPrintable(char byte) bool {
	return 32 <= char && char <= 126
}

func isWordCharacter(char byte) bool {
	if char == ' ' || char == '\n' {
		return true
	}
	char &^= 0x20 // convert to uppercase if lowercase
	return 'A' <= char && char <= 'Z'
}

func replaceNewlines(in []byte) []byte {
	for i := 0; i < len(in); i++ {
		if in[i] == '\n' {
			in[i] = ' '
		}
	}
	return in
}

func Score(in []byte) int {
	var score int

	for _, char := range in {
		if isWordCharacter(char) {
			score += 1
		} else if !isPrintable(char) {
			return 0
		}
	}

	return score
}

func BestKey(in []byte, verbose bool) (byte, int) {
	var n = len(in)
	var out = make([]byte, n)

	var bestScore, score int
	var bestKey, key byte

	bestScore = Score(in)

	// we have already considered key = 0 by scoring in
	for key = 1; key != 0; key++ {
		for i := 0; i < n; i++ {
			out[i] = in[i] ^ key
		}
		score = Score(out)
		if score > bestScore {
			bestScore = score
			bestKey = key
		}
		if verbose && score > 0 {
			fmt.Printf(
				" key: %2x score: %2d out: %s\n",
				key, score, replaceNewlines(out),
			)
		}
	}

	return bestKey, bestScore
}

func Break(in []byte, verbose bool) ([]byte, int) {
	bestKey, bestScore := BestKey(in, verbose)
	for i := 0; i < len(in); i++ {
		in[i] ^= bestKey
	}

	return in, bestScore
}
