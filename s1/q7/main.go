package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s1/q7/ecb"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: q7 KEY\n")
		return
	}
	var key = []byte(os.Args[1])

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	out, err := ecb.Decrypt(buf.Bytes(), key)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return
	}

	os.Stdout.Write(out)
}
