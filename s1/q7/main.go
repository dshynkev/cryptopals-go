package main

import (
	"bufio"
	"fmt"
	"os"

	"cryptopals/common"
	"cryptopals/s1/q7/ecb"
)

func main() {
	var in, key []byte
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Provide a key")
		return
	}
	key = []byte(os.Args[1])

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = append(in, common.B64ToRaw(scanner.Bytes())...)
	}

	out, err := ecb.Decrypt(in, key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Decryption error: %s", err)
		return
	}

	fmt.Println(string(out))
}
