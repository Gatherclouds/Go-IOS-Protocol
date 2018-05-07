package iostdb

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
)

func Sha256(raw []byte) []byte {
	var data = sha256.Sum256(raw)
	return data[:]
}

func Hash160(raw []byte) []byte {
	var data = sha256.Sum256(raw)
	return ripemd160.New().Sum(data[len(data):])
}

func Base58Encode(raw []byte) string {
	return base58.Encode(raw)
}

func Base58Decode(s string) []byte {
	return base58.Decode(s)
}

func ToHex(data []byte) string {
	return hex.EncodeToString(data)
}

func ParseHex(s string) []byte {
	d, err := hex.DecodeString(s)
	if err != nil {
		println(err)
		return nil
	}
	return d
}




