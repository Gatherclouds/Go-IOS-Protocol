package rlp

import (
	"errors"
	"io"
	"bytes"
)

var (
	errNoPointer     = errors.New("rlp: interface given to Decode must be a pointer")
	errDecodeIntoNil = errors.New("rlp: pointer given to Decode must not be nil")
)

type Decoder interface {
	DecodeRLP(*Stream) error
}

func Decode(r io.Reader, val interface{}) error {
	// TODO: this could use a Stream from a pool.
	return NewStream(r, 0).Decode(val)
}

func DecodeBytes(b []byte, val interface{}) error {
	// TODO: this could use a Stream from a pool.
	r := bytes.NewReader(b)
	if err := NewStream(r, uint64(len(b))).Decode(val); err != nil {
		return err
	}
	if r.Len() > 0 {
		return ErrMoreThanOneValue
	}
	return nil
}