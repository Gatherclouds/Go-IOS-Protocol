package iostdb

import "fmt"

func (bc *BlockChainImpl) Get(layer int) (*Block, error) {

	if layer < 0 || layer >= bc.length {
		return nil, fmt.Errorf("index exceed")
	}
	
	var block Block
	block.Decode(blk)
	return &block, nil
}





