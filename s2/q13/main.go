package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s2/q13/oracle"
)

func main() {
	re := oracle.NewEncryptor()

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	out := re.Encrypt(buf.Bytes())
	os.Stdout.Write(out)
}
