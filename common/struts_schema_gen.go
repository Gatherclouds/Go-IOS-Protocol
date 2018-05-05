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

func (d *SignatureRaw) Size() (s uint64) {

	{
		l := uint64(len(d.Sig))

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
	{
		l := uint64(len(d.Pubkey))

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
	s += 1
	return
}

func (d *SignatureRaw) Marshal(buf []byte) ([]byte, error) {
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

		buf[0+0] = byte(d.Algorithm >> 0)

	}
	{
		l := uint64(len(d.Sig))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+1] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+1] = byte(t)
			i++

		}
		copy(buf[i+1:], d.Sig)
		i += l
	}
	{
		l := uint64(len(d.Pubkey))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+1] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+1] = byte(t)
			i++

		}
		copy(buf[i+1:], d.Pubkey)
		i += l
	}
	return buf[:i+1], nil
}
