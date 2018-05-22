package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

type Member struct {
	ID     string
	Pubkey []byte
	Seckey []byte
}

func NewMember(seckey []byte) (Member, error) {
	var m Member
	if seckey == nil {
		seckey = randomSeckey()
	}
	
	m.Seckey = seckey
	m.Pubkey = makePubkey(seckey)
	m.ID = GetIdByPubkey(m.Pubkey)
	return m, nil
}



