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
func (d *UTXO) Size() (s uint64) {

	{
		l := uint64(len(d.BirthTxHash))

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
		l := uint64(len(d.Script))

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
	s += 8
	return
}
func (d *UTXO) Marshal(buf []byte) ([]byte, error) {
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
		l := uint64(len(d.BirthTxHash))

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
		copy(buf[i+0:], d.BirthTxHash)
		i += l
	}
	{

		buf[i+0+0] = byte(d.Value >> 0)

		buf[i+1+0] = byte(d.Value >> 8)

		buf[i+2+0] = byte(d.Value >> 16)

		buf[i+3+0] = byte(d.Value >> 24)

		buf[i+4+0] = byte(d.Value >> 32)

		buf[i+5+0] = byte(d.Value >> 40)

		buf[i+6+0] = byte(d.Value >> 48)

		buf[i+7+0] = byte(d.Value >> 56)

	}
	{
		l := uint64(len(d.Script))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+8] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+8] = byte(t)
			i++

		}
		copy(buf[i+8:], d.Script)
		i += l
	}
	return buf[:i+8], nil
}
func (d *UTXO) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.BirthTxHash)) >= l {
			d.BirthTxHash = d.BirthTxHash[:l]
		} else {
			d.BirthTxHash = make([]byte, l)
		}
		copy(d.BirthTxHash, buf[i+0:])
		i += l
	}
	{

		d.Value = 0 | (int64(buf[i+0+0]) << 0) | (int64(buf[i+1+0]) << 8) | (int64(buf[i+2+0]) << 16) | (int64(buf[i+3+0]) << 24) | (int64(buf[i+4+0]) << 32) | (int64(buf[i+5+0]) << 40) | (int64(buf[i+6+0]) << 48) | (int64(buf[i+7+0]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+8] & 0x7F)
			for buf[i+8]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+8]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Script = string(buf[i+8 : i+8+l])
		i += l
	}
	return i + 8, nil
}

type TxInput struct {
	TxHash       []byte
	UnlockScript string
	UTXOHash     []byte
}