package core

import (
	"bytes"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type UTXOPool interface {
	Add(state UTXO) error
	Find(stateHash []byte) (UTXO, error)
	Del(StateHash []byte) error
	Transact(block *Block) error
	Flush() error
	Copy() UTXOPool
}