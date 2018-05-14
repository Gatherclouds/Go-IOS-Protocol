package account

import (
	"github.com/LoCCS/bliss/params"
	"github.com/LoCCS/bliss"
	"github.com/LoCCS/bliss/sampler"
)

const (
	AddressLength = 52 // or 43 if use base64
	BlissVersion  = params.BLISS_B_4
)

type Address struct {
	pk  bliss.PublicKey
	txt [AddressLength]byte
}

func (addr *Address) ToString() string { return string(addr.txt[:]) }

func newSeed(str string) []uint8 {
	str_len := uint32(len(str))
	var seed_size uint32
	if str_len < sampler.SHA_512_DIGEST_LENGTH {
		seed_size = sampler.SHA_512_DIGEST_LENGTH
	} else {
		seed_size = str_len
	}

	seed := make([]uint8, seed_size)
	copy(seed[:str_len], ([]uint8)(str))
	return seed
}

func newPrivateKey(str string) (*sampler.Entropy, *bliss.PrivateKey, error) {
	seed := newSeed(str)
	entropy, err := sampler.NewEntropy(seed)
	if err != nil {
		return nil, nil, err
	}

	sk, err := bliss.GeneratePrivateKey(BlissVersion, entropy)
	if err != nil {
		return nil, nil, err
	} else {
		return entropy, sk, nil
	}
}

