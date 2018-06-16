package rlp

import (
	"io"
	"sync"
)

var (
	EmptyString = []byte{0x80}
	EmptyList   = []byte{0xC0}
)

type Encoder interface {
	EncodeRLP(io.Writer) error
}

func Encode(w io.Writer, val interface{}) error {
	if outer, ok := w.(*encbuf); ok {
		return outer.encode(val)
	}
	eb := encbufPool.Get().(*encbuf)
	defer encbufPool.Put(eb)
	eb.reset()
	if err := eb.encode(val); err != nil {
		return err
	}
	return eb.toWriter(w)
}

func EncodeToBytes(val interface{}) ([]byte, error) {
	eb := encbufPool.Get().(*encbuf)
	defer encbufPool.Put(eb)
	eb.reset()
	if err := eb.encode(val); err != nil {
		return nil, err
	}
	return eb.toBytes(), nil
}

func EncodeToReader(val interface{}) (size int, r io.Reader, err error) {
	eb := encbufPool.Get().(*encbuf)
	eb.reset()
	if err := eb.encode(val); err != nil {
		return 0, nil, err
	}
	return eb.size(), &encReader{buf: eb}, nil
}

type encbuf struct {
	str     []byte
	lheads  []*listhead
	lhsize  int
	sizebuf []byte
}

type listhead struct {
	offset int
	size   int
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

// puthead writes a list or string header to buf.
// buf must be at least 9 bytes long.
func puthead(buf []byte, smalltag, largetag byte, size uint64) int {
	if size < 56 {
		buf[0] = smalltag + byte(size)
		return 1
	} else {
		sizesize := putint(buf[1:], size)
		buf[0] = largetag + byte(sizesize)
		return sizesize + 1
	}
}

// encbufs are pooled.
var encbufPool = sync.Pool{
	New: func() interface{} { return &encbuf{sizebuf: make([]byte, 9)} },
}

func (w *encbuf) reset() {
	w.lhsize = 0
	if w.str != nil {
		w.str = w.str[:0]
	}
	if w.lheads != nil {
		w.lheads = w.lheads[:0]
	}
}

// encbuf implements io.Writer so it can be passed it into EncodeRLP.
func (w *encbuf) Write(b []byte) (int, error) {
	w.str = append(w.str, b...)
	return len(b), nil
}
