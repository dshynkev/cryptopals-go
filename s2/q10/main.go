package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s2/q10/cbc"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: q10 KEY\n")
		return
	}

	var key = []byte(os.Args[1])

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	if buf.Len() < cbc.BlockSize {
		os.Stderr.WriteString("Ciphertext must contain IV\n")
		return
	}

	var iv = buf.Bytes()[:cbc.BlockSize]
	var ciphertext = buf.Bytes()[cbc.BlockSize:]
	var out, err = cbc.Decrypt(ciphertext, key, iv)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return
	}

	var rd = bytes.NewReader(out)
	io.Copy(os.Stdout, rd)
}
