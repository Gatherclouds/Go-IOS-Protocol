package iosbase

import (
	"fmt"
	"math/rand"
	"time"
)

type Member struct {
	ID       string
	userName string
	PubKey   []byte
	SecKey   []byte
}

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
	m.PubKey = CalcPubkey(secKey)
	m.ID = Base58Encode(m.PubKey)
	return m, nil
}
