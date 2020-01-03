package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s1/q6/repeated"
)

func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	out, err := repeated.Break(buf.Bytes(), true)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return
	}
	os.Stdout.Write(out)
}
