package ecb

import (
  "io/ioutil"
	"testing"

	"cryptopals/common"
)

var key = []byte("YELLOW SUBMARINE")

func TestDecrypt(t *testing.T) {
  in, _ := ioutil.ReadFile("in.txt")
  want, _ := ioutil.ReadFile("out.txt")

  in = common.Expunge(in, '\n')

  got, err := Decrypt(common.B64ToRaw(in), key)
  if err != nil {
    t.Fatalf("got err; want %s", got)
  }

  common.Test(t, want, got)
}
