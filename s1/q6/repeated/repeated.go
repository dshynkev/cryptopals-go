package repeated

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"math/bits"

	"cryptopals/s1/q3/sbxor"
)

var CouldNotBreak = errors.New("could not break code")

func EditDistance(lhs, rhs []byte) int {
	var mask uint8
	var distance int

	for i := 0; i < len(lhs); i++ {
		mask = uint8(lhs[i] ^ rhs[i])
		distance += bits.OnesCount8(mask)
	}

	return distance
}

const (
	BlockSamples = 8
	MinKeylen    = 2
	MaxKeylen    = 40
)

func BestKeylen(in []byte, verbose bool) int {
	maxKeylen := len(in) / BlockSamples
	if maxKeylen > MaxKeylen {
		maxKeylen = MaxKeylen
	}

	var (
		distance     float64
		bestDistance float64 = math.Inf(+1)
		keylen       int
		bestKeylen   int = 0
	)

	for keylen = MinKeylen; keylen < maxKeylen; keylen++ {
		distance = 0

		for i := 0; i < BlockSamples; i += 2 {
			distance += float64(
				EditDistance(
					in[(i+0)*keylen:(i+1)*keylen],
					in[(i+1)*keylen:(i+2)*keylen],
				),
			)
		}
		distance /= float64(keylen)

		if distance < bestDistance {
			bestDistance = distance
			bestKeylen = keylen
		}

		if verbose {
			fmt.Printf("  keylen: %2d distance: %7.4f\n", keylen, distance)
		}
	}

	if verbose {
		fmt.Printf(" chose keylen: %2d\n", bestKeylen)
	}
	return bestKeylen
}

func split(in []byte, keylen int) [][]byte {
	var padding = keylen - len(in)%keylen
	in = append(in, bytes.Repeat([]byte{0}, padding)...)

	var n = (len(in) + keylen - 1) / keylen
	var chunks = make([][]byte, keylen)

	for i := 0; i < keylen; i++ {
		chunks[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			chunks[i][j] = in[i+j*keylen]
		}
	}

	return chunks
}

func join(chunks [][]byte) []byte {
	var (
		keylen = len(chunks)
		n      = len(chunks[0])
	)
	var united = make([]byte, n*keylen)

	for i := 0; i < keylen; i++ {
		for j := 0; j < n; j++ {
			united[i+j*keylen] = chunks[i][j]
		}
	}

	return united
}

func Break(in []byte, verbose bool) ([]byte, error) {
	bestKeylen := BestKeylen(in, verbose)
	if bestKeylen == 0 {
		return nil, CouldNotBreak
	}

	var score int
	chunks := split(in, bestKeylen)
	for i, chunk := range chunks {
		chunks[i], score = sbxor.Break(chunk, verbose)
		if score == 0 {
			return nil, CouldNotBreak
		}
	}
	out := join(chunks)[:len(in)]

	return out, nil
}
