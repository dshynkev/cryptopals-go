package main

import (
	"bufio"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q8/detect"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	in := common.HexToRaw(scanner.Bytes())

	if detect.IsEcb(in) {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
