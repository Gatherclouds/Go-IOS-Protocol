package block

import (
	_ "github.com/gomodule/redigo/redis"
	"Go-IOS-Protocol/db"
	"Go-IOS-Protocol/core/tx"
)

var (
	blockLength = []byte("BlockLength") //blockLength -> length of ChainImpl

	blockNumberPrefix = []byte("n") //blockNumberPrefix + block number -> block hash
	blockPrefix       = []byte("H") //blockHashPrefix + block hash -> block data
)

// ChainImpl 是已经确定block chain的结构体
type ChainImpl struct {
	db     db.Database
	length uint64
	tx     tx.TxPool
}

//func (bc *BlockChainImpl) Get(layer int) (*Block, error) {
//
//	if layer < 0 || layer >= bc.length {
//		return nil, fmt.Errorf("index exceed")
//	}
//
//	headHash, err := redis.Bytes(bc.redis.Do("LINDEX", IndexKey, layer))
//	if err != nil {
//		return nil, err
//	}
//
//	blk, err := bc.db.Get(headHash)
//	if err != nil {
//		return nil, err
//	}
//	var block Block
//	block.Decode(blk)
//	return &block, nil
//}
//
//func (bc *BlockChainImpl) Push(block *Block) error {
//	err := bc.db.Put(block.HeadHash(), block.Encode())
//	if err != nil {
//		return err
//	}
//
//	_, err = bc.redis.Do("RPUSH", IndexKey, block.HeadHash())
//	bc.length++
//	return nil
//}
//
//func (bc *BlockChainImpl) Length() int {
//	return bc.length
//}
//
//func (bc *BlockChainImpl) Top() *Block {
//	blk, err := bc.Get(bc.length - 1)
//	if err != nil {
//		panic(err)
//	}
//	return blk
//}
//func (bc *BlockChainImpl) Init() error {
//	var err error
//	bc.db, err = iostdb.NewLDBDatabase(DBPath, 1, 1)
//	if err != nil {
//		return err
//	}
//
//	bc.redis, err = redis.Dial(Conn, DBAddr)
//	if err != nil {
//		return err
//	}
//
//	len, err := redis.Int(bc.redis.Do("llen", "BC_index"))
//	if err != nil {
//		return err
//	}
//	bc.length = len
//	return nil
//}
//
//func (bc *BlockChainImpl) Close() error {
//	bc.db.Close()
//	return nil
//}