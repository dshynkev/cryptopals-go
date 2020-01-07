package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s3/q17/oracle"
)

func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	re := oracle.NewEncryptor()
	valid := re.CheckPadding(buf.Bytes())

	if valid {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
