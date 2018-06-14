package rlp

import "io"

var (
	EmptyString = []byte{0x80}
	EmptyList   = []byte{0xC0}
)

type Encoder interface {
	EncodeRLP(io.Writer) error
}

// encode writes head to the given buffer, which must be at least
// 9 bytes long. It returns the encoded bytes.
func (head *listhead) encode(buf []byte) []byte {
	return buf[:puthead(buf, 0xC0, 0xF7, uint64(head.size))]
}

// headsize returns the size of a list or string header
// for a value of the given size.
func headsize(size uint64) int {
	if size < 56 {
		return 1
	}
	return 1 + intsize(size)
}

