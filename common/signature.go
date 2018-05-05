package common

import "fmt"

type SignAlgorithm uint8

const (
	Secp256k1 SignAlgorithm = iota
)

type Signature struct {
	Algorithm SignAlgorithm

	Sig    []byte
	Pubkey []byte
}

func Sign(algo SignAlgorithm, info, privkey []byte) (Signature, error) {
	s := Signature{}
	s.Algorithm = algo
	switch algo {
	case Secp256k1:
		s.Pubkey = CalcPubkeyInSecp256k1(privkey)
		s.Sig = SignInSecp256k1(info, privkey)
		return s, nil
	}
	return s, fmt.Errorf("algorithm not exist")
}

func VerifySignature(info []byte, s Signature) bool {
	switch s.Algorithm {
	case Secp256k1:
		return VerifySignInSecp256k1(info, s.Pubkey, s.Sig)
	}
	return false
}

