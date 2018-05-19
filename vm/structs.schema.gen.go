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

func (d *contractInfoRaw) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Language))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Language)
		i += l
	}
	{

		buf[i+0+0] = byte(d.Version >> 0)

	}
	{

		buf[i+0+1] = byte(d.GasLimit >> 0)

		buf[i+1+1] = byte(d.GasLimit >> 8)

		buf[i+2+1] = byte(d.GasLimit >> 16)

		buf[i+3+1] = byte(d.GasLimit >> 24)

		buf[i+4+1] = byte(d.GasLimit >> 32)

		buf[i+5+1] = byte(d.GasLimit >> 40)

		buf[i+6+1] = byte(d.GasLimit >> 48)

		buf[i+7+1] = byte(d.GasLimit >> 56)

	}

	return buf[:i+17], nil
}
