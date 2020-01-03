package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s2/q14/oracle"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: q14 SECRET\n")
		return
	}

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	re := oracle.NewEncryptor([]byte(os.Args[1]))
	out := re.Encrypt(buf.Bytes())

	os.Stdout.Write(out)
}
