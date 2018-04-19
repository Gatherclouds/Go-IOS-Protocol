package core

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/iost-official/prototype/iostdb"
)

//go:generate mockgen -destination mocks/mock_blockchain.go -package core_mock -source blockchain.go -imports .=github.com/iost-official/prototype/core

// Block chain
type BlockChain interface {
	Get(layer int) (*Block, error)
	Push(block *Block) error
	Length() int
	Top() *Block

	FindTx(txHash []byte) (Tx, error)
}

type BlockChainImpl struct {
	db     *iostdb.LDBDatabase
	redis  redis.Conn
	length int
}

const (
	DBPath   = "savedata/"
	IndexKey = "block_chain_index"
)
