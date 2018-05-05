package common

import (
	"unsafe"
	"io"
	"time"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type SignatureRaw struct {
	Algorithm int8
	Sig       []byte
	Pubkey    []byte
}
