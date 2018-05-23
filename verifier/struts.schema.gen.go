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

func (d *VerifyLogRaw) Size() (s uint64) {

	{
		l := uint64(len(d.List))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

	}
	return
}
