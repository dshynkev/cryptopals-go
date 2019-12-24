package main

import (
	"bufio"
	"os"

	"cryptopals/s2/q13/oracle"
)

func main() {
	re := oracle.NewEncryptor()

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	out := re.Encrypt(scanner.Bytes())
	os.Stdout.Write(out)
}
