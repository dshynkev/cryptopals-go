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

	out := convert.HexToB64(in)
	fmt.Println(string(out))
}
