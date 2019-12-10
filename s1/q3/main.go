package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q3/sbxor"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	in := common.HexToRaw(scanner.Bytes())

	out, score := sbxor.Break(in, true)
	if score == 0 {
		os.Exit(-1)
	}
	fmt.Println(string(out))
}
