package drp

import (
	"fmt"
	"math/big"
)

type ShareId int

type EncryptedShare struct {
	sid          ShareId
	encryptedVal Point
}

type DecryptedShare struct {
	sid          ShareId
	decryptedVal Point
}


