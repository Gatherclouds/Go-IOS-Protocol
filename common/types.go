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
