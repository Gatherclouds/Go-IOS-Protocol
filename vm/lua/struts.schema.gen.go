package lua

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

type contractRaw struct {
	info    []byte
	code    []byte
	methods []methodRaw
}

