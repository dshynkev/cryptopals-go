package convert

import "cryptopals/common"

func HexToB64(in []byte) []byte {
	return common.RawToB64(common.HexToRaw(in))
}
