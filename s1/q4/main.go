package main

import (
	"bufio"
	"os"

	"cryptopals/common/convert"
	"cryptopals/s1/q4/detect"
)

func main() {
	var lines [][]byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, convert.HexToRaw(scanner.Bytes()))
	}

	out, _, score := detect.SBXor(lines, true)
	if score == 0 {
		os.Exit(-1)
	}
	os.Stdout.Write(out)
}
