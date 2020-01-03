package main

import (
	"bytes"
	"io"
	"os"

	"cryptopals/s1/q1/convert"
)

func main() {
	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	out := convert.HexToB64(buf.Bytes())
	os.Stdout.Write(out)
}
