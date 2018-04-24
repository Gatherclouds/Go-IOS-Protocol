package iostdb

import (
	"github.com/gomodule/redigo/redis"
)

//go:generate mockgen -destination mocks/mock_statepool.go -package iosbase_mock -source statepool.go -imports .=github.com/iost-official/Go-IOS-Protocol/iosbase

/*
Current states of system
*/
type StatePool interface {
	Init() error
	Close() error
	Add(state State) error
	Find(stateHash []byte) (State, error)
	Del(StateHash []byte) error
	Transact(block *Block) error
}

