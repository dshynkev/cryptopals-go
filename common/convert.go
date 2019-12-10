package common

import (
	"encoding/base64"
	"encoding/hex"
)

var b64 = base64.StdEncoding

func HexToRaw(in []byte) []byte {
	var out = make([]byte, hex.DecodedLen(len(in)))

	_, err := hex.Decode(out, in)
	if err != nil {
		panic(err)
	}

	return out
}

func RawToHex(in []byte) []byte {
	var out = make([]byte, hex.EncodedLen(len(in)))

	hex.Encode(out, in)

	return out
}

func B64ToRaw(in []byte) []byte {
	var out = make([]byte, b64.DecodedLen(len(in)))

	_, err := b64.Decode(out, in)
	if err != nil {
		panic(err)
	}

	return out
}

func RawToB64(in []byte) []byte {
	var out = make([]byte, b64.EncodedLen(len(in)))

	b64.Encode(out, in)

	return out
}
