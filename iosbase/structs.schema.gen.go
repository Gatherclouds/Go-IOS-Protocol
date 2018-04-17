package iosbase

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type State struct {
	BirthTextHash []byte
	Value       int64
	Script      string
}

