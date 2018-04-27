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





