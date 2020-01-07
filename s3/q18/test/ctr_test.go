package ctr_test

import (
	"testing"

	"cryptopals/common/convert"
	"cryptopals/common/test"
	"cryptopals/s3/q18/ctr"
)

func TestDecrypt(t *testing.T) {
	want := []byte("Yo, VIP Let's kick it Ice, Ice, baby Ice, Ice, baby ")
	in := []byte("L77na/nrFsKvynd6HzOoG7GHTLXsTVu9qvY/2syLXzhPweyyMTJULu/6/kXX0KSvoOLSFQ==")
	key := []byte("YELLOW SUBMARINE")
	nonce := make([]byte, ctr.NonceSize)

	got, err := ctr.Decrypt(convert.B64ToRaw(in), key, nonce)
	test.Test(t, nil, err)
	test.Test(t, want, got)
}
