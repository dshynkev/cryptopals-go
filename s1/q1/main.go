package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/s1/q1/convert"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	in := scanner.Bytes()

	out, err := convert.HexToB64(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		return
	}
	fmt.Println(string(out))
}
