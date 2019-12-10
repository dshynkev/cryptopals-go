package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q4/detect"
)

func main() {
	var lines [][]byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, common.HexToRaw(scanner.Bytes()))
	}

	out, _, score := detect.SBXor(lines, true)
	if score == 0 {
		os.Exit(-1)
	}
	fmt.Print(string(out))
}
