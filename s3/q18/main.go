package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s3/q18/ctr"
)

func main() {
	if len(os.Args) != 2 {
		os.Stderr.WriteString("Usage: q18 KEY\n")
		return
	}

	var key = []byte(os.Args[1])
	var nonce = make([]byte, ctr.NonceSize)

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	var out, err = ctr.Encrypt(buf.Bytes(), key, nonce)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return
	}

	os.Stdout.Write(out)
}
