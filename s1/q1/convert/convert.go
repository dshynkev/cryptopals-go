package convert

import "cryptopals/common/convert"

func HexToB64(in []byte) []byte {
	return convert.RawToB64(convert.HexToRaw(in))
}
