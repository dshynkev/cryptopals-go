package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q5/repeated"
)

func main() {
	var in, key, out []byte

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Provide a key")
		return
	}
	key = []byte(os.Args[1])

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
    return
  }
  in = scanner.Bytes()
  out = common.RawToHex(repeated.Xor(in, key))
  fmt.Println(string(out))
}
