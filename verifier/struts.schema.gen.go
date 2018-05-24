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

		for k0 := range d.List {

			{
				l := uint64(len(d.List[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}

				s += 4 * l

			}
		}
	}
	return
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

		for k0 := range d.List {

			{
				l := uint64(len(d.List[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}

				s += 4 * l

			}

		}

	}
	return
}

func (d *VerifyLogRaw) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.List))

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
		for k0 := range d.List {

			{
				l := uint64(len(d.List[k0]))

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
				
			}

		}
	}
	return buf[:i+0], nil
}
