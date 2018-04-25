package iostdb

import (
	"bytes"
	"encoding/binary"
)

func NewBinary() Binary {
	var b Binary
	b.bytes = []byte{}
	b.length = 0
	return b
}

func BinaryFromBase58(s string) Binary {
	bits := Base58Decode(s)
	return Binary{bits, len(bits)}
}

func (bin *Binary) ToBase58() string {
	return Base58Encode(bin.bytes)
}

func (bin *Binary) PutUInt(u uint32) *Binary {
	var bBuf = bytes.NewBuffer([]byte{})
	binary.Write(bBuf, binary.BigEndian, u)
	bin.bytes = append(bin.bytes, bBuf.Bytes()...)
	bin.length += 4
	return bin
}












