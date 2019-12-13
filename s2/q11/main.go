package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s2/q11/oracle"
)

func main() {
	var buf bytes.Buffer

	io.Copy(&buf, os.Stdin)

	re := oracle.NewRandomEncryptor()
	out := re.Ask(buf.Bytes())

	os.Stdout.Write(out)
}
