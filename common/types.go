package common

const (
	HashLength    = 32
	AddressLength = 20
)

type Hash [HashLength]byte

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func HexToHash(s string) Hash		{ return BytesToHash(FromHex(s)) }
func (h Hash) Bytes() []byte { return h[:] }

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}
