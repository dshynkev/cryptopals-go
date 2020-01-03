package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/common/pkcs7"
)

const blockSize = 16

func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	out := pkcs7.Pad(buf.Bytes(), blockSize)

	os.Stdout.Write(out)
}
