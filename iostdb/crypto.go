package iostdb

import (
	"crypto/sha256"
)

func Sha256(raw []byte) []byte {
	var data = sha256.Sum256(raw)
	return data[:]
}








