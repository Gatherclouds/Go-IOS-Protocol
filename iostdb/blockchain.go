package iostdb

import (
	"github.com/gomodule/redigo/redis"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

//go:generate mockgen -destination mocks/mock_blockchain.go -package iosbase_mock -source blockchain.go -imports .=github.com/iost-official/Go-IOS-Protocol/iosbase

// Block chain
type BlockChain interface {
	Get(layer int) (*Block, error)
	Push(block *Block) error
	Length() int
	Top() *Block
	Init() error
	Close() error
}


