package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"cryptopals/s2/q12/oracle"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s SECRET\n", os.Args[0])
		return
	}

	var buf bytes.Buffer
	io.Copy(&buf, os.Stdin)

	re := oracle.NewEncryptor([]byte(os.Args[1]))
	out := re.Ask(buf.Bytes())

	os.Stdout.Write(out)
}
