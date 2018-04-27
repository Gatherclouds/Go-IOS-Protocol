package iostdb

import "fmt"

func (bc *BlockChainImpl) Get(layer int) (*Block, error) {

	if layer < 0 || layer >= bc.length {
		return nil, fmt.Errorf("index exceed")
	}

	headHash, err := redis.Bytes(bc.redis.Do("LINDEX", IndexKey, layer))
	if err != nil {
		return nil, err
	}

	blk, err := bc.db.Get(headHash, &opt.ReadOptions{})
	if err != nil {
		return nil, err
	}
	var block Block
	block.Decode(blk)
	return &block, nil
}

func (bc *BlockChainImpl) Push(block *Block) error {
	err := bc.db.Put(block.HeadHash(), block.Encode(), &opt.WriteOptions{})
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



