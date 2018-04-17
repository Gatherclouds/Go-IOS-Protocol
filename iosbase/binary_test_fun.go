package iosbase

import (
	"math/rand"
	"time"
	"fmt"
)

func NewMember(secKey []byte) (Member, error) {
	var m Member
	if secKey == nil {
		rand.Seed(time.Now().UnixNano())
		bin := NewBinary()
		for i := 0; i < 4; i++ {
			bin.PutULong(rand.Uint64())
		}
		secKey = bin.bytes
	}
	if len(secKey) != 32 {
		return Member{}, fmt.Errorf("seckey length error")
	}

	m.SecKey = secKey
	m.PubKey = CalcPubKey(secKey)
	m.ID = Base58Encode(m.PubKey)
	return m, nil
}
