package vm

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

type contractInfoRaw struct {
	Language string
	Version  int8
	GasLimit int64
	Price    float64
}

