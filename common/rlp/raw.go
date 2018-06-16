package rlp

import "reflect"

type RawValue []byte

var rawValueType = reflect.TypeOf(RawValue{})

func ListSize(contentSize uint64) uint64 {
	return uint64(headsize(contentSize)) + contentSize
}

