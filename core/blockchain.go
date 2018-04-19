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
func (bc *BlockChainImpl) Get(layer int) (*Block, error) {

	if layer < 0 || layer >= bc.length {
		return nil, fmt.Errorf("index exceed")
	}

	headHash, err := redis.Bytes(bc.redis.Do("LINDEX", IndexKey, layer))
	if err != nil {
		return nil, err
	}

	blk, err := bc.db.Get(headHash)
	if err != nil {
		return nil, err
	}
	var block Block
	block.Decode(blk)
	return &block, nil
}

func (bc *BlockChainImpl) Push(block *Block) error {
	err := bc.db.Put(block.HeadHash(), block.Encode())
	if err != nil {
		return err
	}

	_, err = bc.redis.Do("RPUSH", IndexKey, block.HeadHash())
	bc.length++
	return nil
}

func (bc *BlockChainImpl) Length() int {
	return bc.length
}

func (bc *BlockChainImpl) Top() *Block {
	blk, err := bc.Get(bc.length - 1)
	if err != nil {
		panic(err)
	}
	return blk
}