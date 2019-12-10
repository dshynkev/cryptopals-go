package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q6/repeated"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	in := common.HexToRaw(scanner.Bytes())

	out, err := repeated.Break(in, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		return
	}
	fmt.Println(string(out))
}
