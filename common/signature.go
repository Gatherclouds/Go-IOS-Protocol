package common

type SignAlgorithm uint8

const (
	Secp256k1 SignAlgorithm = iota
)

type Signature struct {
	Algorithm SignAlgorithm

	Sig    []byte
	Pubkey []byte
}
