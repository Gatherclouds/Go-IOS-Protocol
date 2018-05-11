package pow

import "github.com/ethereum/go-ethereum/core"

type CachedBlockChain struct {
	core.BlockChain
	cachedBlock []*core.Block
}

func NewCBC(chain core.BlockChain) CachedBlockChain {
	return CachedBlockChain{
		BlockChain:  chain,
		cachedBlock: make([]*core.Block, 0),
	}
}

