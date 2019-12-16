package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/commmon/pkcs7"
)

// This reimplements pkcs.Pad for efficiency reasons
func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	padding := pkcs7.BlockSize - buf.Len()%pkcs7.BlockSize
	for i := 0; i < padding; i++ {
		buf.WriteByte(byte(padding))
	}

	io.Copy(os.Stdout, &buf)
}
