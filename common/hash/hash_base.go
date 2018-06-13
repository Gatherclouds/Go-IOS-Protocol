package hash

const (
	HashLength     = 32
	LongHashLength = 64
)

type Hash [HashLength]byte
type LongHash [LongHashLength]byte

func NewHash(bs []byte) Hash {
	bLen := len(bs)
	var h Hash

	if bLen < HashLength {
		for i, b := range bs {
			h[HashLength-bLen+i] = b
		}
	} else {
		for i := range h {
			h[i] = bs[i]
		}
	}

	return h
}
