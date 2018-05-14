package account

import (
	"github.com/LoCCS/bliss/params"
	"github.com/LoCCS/bliss"
)

const (
	AddressLength = 52 // or 43 if use base64
	BlissVersion  = params.BLISS_B_4
)

type Address struct {
	pk  bliss.PublicKey
	txt [AddressLength]byte
}

func (addr *Address) ToString() string { return string(addr.txt[:]) }


