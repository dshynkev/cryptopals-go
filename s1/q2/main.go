package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q2/xor"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	in1 := common.HexToRaw(scanner.Bytes())
	if !scanner.Scan() {
		return
	}
	in2 := common.HexToRaw(scanner.Bytes())

	out := common.RawToHex(xor.Xor(in1, in2))
	fmt.Println(string(out))
}
