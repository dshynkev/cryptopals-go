package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"cryptopals/s2/q10/cbc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s KEY\n", os.Args[0])
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
