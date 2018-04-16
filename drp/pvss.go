package drp

import (
	_ "fmt"
	_ "math/big"
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


