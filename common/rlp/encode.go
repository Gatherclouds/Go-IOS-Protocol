package rlp

import "io"

var (
	EmptyString = []byte{0x80}
	EmptyList   = []byte{0xC0}
)

type Encoder interface {
	EncodeRLP(io.Writer) error
}

