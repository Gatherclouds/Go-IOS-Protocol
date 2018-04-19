package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/iost-official/prototype/common"
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
	if len(seckey) != 32 {
		return Member{}, fmt.Errorf("seckey length error")
	}

	m.Seckey = seckey
	m.Pubkey = makePubkey(seckey)
	m.ID = GetIdByPubkey(m.Pubkey)
	return m, nil
}

