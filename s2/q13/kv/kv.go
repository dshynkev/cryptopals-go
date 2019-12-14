package kv

import (
	"cryptopals/common/block"
)

type BidirOracle interface {
	Ask([]byte) []byte
	Tell([]byte) []byte
}

const (
	EmailKeyLength  = len("email=")
	AdminRoleLength = len("admin")
	UserRoleLength  = len("user")
)

func fill(buf []byte, count int, value byte) {
	for i := 0; i < count; i++ {
		buf[i] = value
	}
}

func Break(oracle block.EncryptDecryptOracle) []byte {
	var block = block.GetEcbBlockLayout(oracle)

	var (
		adminOffset        = block.Size - EmailKeyLength
		adminPaddingOffset = adminOffset + AdminRoleLength
		adminPaddingLength = block.Size - AdminRoleLength
		adminPaddingValue  = adminPaddingLength
		adminPayloadLength = adminPaddingOffset + adminPaddingLength
	)

	var scratch = make([]byte, 2*block.Size)
	var payload = make([]byte, (block.Count+1)*block.Size)

	fill(scratch, block.Padding, 0x42)
	copy(scratch[adminOffset:], []byte("admin"))
	fill(scratch[adminPaddingOffset:], adminPaddingLength, byte(adminPaddingValue))

	var response []byte
	// Break the blocks like so: ...[...&role=][user<pad>]
	response = oracle.Encrypt(scratch[:block.Padding+UserRoleLength])
	copy(payload[:block.Count*block.Size], response[:block.Count*block.Size])

	// Break the blocks like so: [email=<pad>][admin<pad>][&...]
	response = oracle.Encrypt(scratch[:adminPayloadLength])
	copy(payload[block.Count*block.Size:], response[block.Size:])

	return oracle.Decrypt(payload)
}
