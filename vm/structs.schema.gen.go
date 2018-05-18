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

func (d *contractInfoRaw) Size() (s uint64) {

	{
		l := uint64(len(d.Language))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 17
	return
}

