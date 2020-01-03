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

	out, err := pkcs7.Unpad(buf.Bytes(), blockSize)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return
	}

	os.Stdout.Write(out)
}
