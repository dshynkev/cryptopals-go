package main

import (
	"bufio"
	"os"

	"cryptopals/common/convert"
	"cryptopals/s1/q2/xor"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	in1 := convert.HexToRaw(scanner.Bytes())
	if !scanner.Scan() {
		return
	}
	in2 := convert.HexToRaw(scanner.Bytes())

	out := xor.Xor(in1, in2)
	os.Stdout.Write(out)
}
