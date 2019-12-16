package main

import (
	"bytes"
	"fmt"
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
            fmt.Fprintln(os.Stderr, err)
        }

        os.Stdout.Write(out)
}
