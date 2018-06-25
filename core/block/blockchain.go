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

// Length 返回已经确定链的长度
func (b *ChainImpl) Length() uint64 {
	return b.length
}

// HasTx 判断tx是否存在于db中
func (b *ChainImpl) HasTx(tx *tx.Tx) (bool, error) {
	return b.tx.Has(tx)
}

// GetTx 通过hash获取tx
func (b *ChainImpl) GetTx(hash []byte) (*tx.Tx, error) {
	return b.tx.Get(hash)
}

// GetBlockByHash 通过区块hash查询块
func (b *ChainImpl) GetBlockByHash(blockHash []byte) *Block {

	block, err := b.db.Get(append(blockPrefix, blockHash...))
	if err != nil {
		return nil
	}
	if len(block) == 0 {
		return nil
	}

	rBlock := new(Block)
	if err := rBlock.Decode(block); err != nil {
		return nil
	}
	return rBlock
}

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