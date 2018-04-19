package core

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

type UTXO struct {
	BirthTxHash []byte
	Value       int64
	Script      string
}
