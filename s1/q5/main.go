package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s1/q5/repeated"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: q5 KEY\n")
		return
	}
	var key = []byte(os.Args[1])

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	var out = repeated.Xor(buf.Bytes(), key)
	os.Stdout.Write(out)
}
