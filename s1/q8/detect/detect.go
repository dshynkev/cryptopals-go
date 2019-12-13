package detect

type empty struct{}

// No need to import crypto/aes just for this
const BlockSize = 16

// IsEcb determines whether the input contains repeated 256-bit blocks.
// An affirmative answer suggests that the input is a AES-ECB ciphertext.
func IsEcb(in []byte) bool {
	var offset, N = 0, len(in) / BlockSize
	var seen = make(map[[BlockSize]byte]empty, N)

	var block [BlockSize]byte
	for i := 0; i < N; i++ {
		copy(block[:], in[offset:(offset+BlockSize)])

		if _, ok := seen[block]; ok {
			return true
		} else {
			seen[block] = empty{}
		}

		offset += BlockSize
	}

	return false
}
