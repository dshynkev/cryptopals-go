package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s1/q3/sbxor"
)

func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	out, score := sbxor.Break(buf.Bytes(), true)
	if score == 0 {
		os.Exit(-1)
	}
	os.Stdout.Write(out)
}
