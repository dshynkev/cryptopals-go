package kv

import (
	"cryptopals/common/edit"
	"cryptopals/common/oracle"
	"cryptopals/common/oracle/ecb"
)

const (
	EmailKeyLength  = len("email=")
	AdminRoleLength = len("admin")
	UserRoleLength  = len("user")
)

func Break(oracle oracle.EncryptDecryptor) []byte {
	layout := ecb.GetBlockLayout(oracle)

	var (
		adminOffset        = layout.BlockSize - EmailKeyLength
		adminPaddingOffset = adminOffset + AdminRoleLength
		adminPaddingLength = layout.BlockSize - AdminRoleLength
		adminPayloadLength = adminPaddingOffset + adminPaddingLength
	)

	var scratch = make([]byte, 2*layout.BlockSize)
	var payload = make([]byte, (layout.BlockCount+1)*layout.BlockSize)

	edit.Fill(scratch[:layout.Padding], 0x42)
	copy(scratch[adminOffset:], []byte("admin"))
	edit.Fill(
		scratch[adminPaddingOffset:adminPaddingOffset+adminPaddingLength],
		byte(adminPaddingLength),
	)

	var response []byte
	// Break the blocks like so: ...[...&role=][user<pad>]
	response = oracle.Encrypt(scratch[:layout.Padding+UserRoleLength])
	copy(
		payload[:layout.BlockCount*layout.BlockSize],
		response[:layout.BlockCount*layout.BlockSize],
	)

	// Break the blocks like so: [email=<pad>][admin<pad>][&...]
	response = oracle.Encrypt(scratch[:adminPayloadLength])
	copy(payload[layout.BlockCount*layout.BlockSize:], response[layout.BlockSize:])

	return oracle.Decrypt(payload)
}
