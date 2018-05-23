package verifier

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

type VerifyLogRaw struct {
	List [][]int32
}
