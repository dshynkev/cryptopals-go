package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s1/q8/detect"
)

func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	if detect.IsEcb(buf.Bytes()) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
