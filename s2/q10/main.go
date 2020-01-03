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
	var iv = make([]byte, cbc.BlockSize)

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	var out, _ = cbc.Decrypt(buf.Bytes(), key, iv)

	var rd = bytes.NewReader(out)
	io.Copy(os.Stdout, rd)
}
