package ecb

import "crypto/aes"

func Unpad(in []byte) []byte {
    N := len(in)
    padding := int(in[N-1])
    return in[:N-padding]
}

func Decrypt(in, key []byte) ([]byte, error) {
  out := make([]byte, len(in))

  cipher, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }

  offset, n := 0, len(in) / aes.BlockSize
  for i := 0; i < n; i++ {
    cipher.Decrypt(
      out[offset:(offset + aes.BlockSize)], in[offset:(offset + aes.BlockSize)],
    )
    offset += aes.BlockSize
  }

  return Unpad(out), nil
}
