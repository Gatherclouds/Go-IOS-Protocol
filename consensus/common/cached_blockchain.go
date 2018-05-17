package common

import "github.com/ethereum/go-ethereum/core"

type CachedBlockChain struct {
	block.Chain
	block        *block.Block
	cachedLength int
	parent       *CachedBlockChain
	depth        int
	// DPoS中使用，记录该节点被最多几个witness确认
	confirmed int
}

func NewCBC(chain block.Chain) CachedBlockChain {
	return CachedBlockChain{
		Chain:        chain,
		block:        nil,
		parent:       nil,
		cachedLength: 0,
		depth:        0,
		confirmed:    0,
	}
}

//func (c *CachedBlockChain) Get(layer int) (*core.Block, error) {
//	if layer < 0 || layer >= c.BlockChain.Length()+len(c.cachedBlock) {
//		return nil, fmt.Errorf("overflow")
//	}
//	if layer < c.BlockChain.Length() {
//		return c.BlockChain.Get(layer)
//	}
//	return c.cachedBlock[layer-c.BlockChain.Length()], nil
//}
func (c *CachedBlockChain) Push(block *core.Block) error {
	c.cachedBlock = append(c.cachedBlock, block)
	return nil
}

func (c *CachedBlockChain) Length() int {
	return c.BlockChain.Length() + len(c.cachedBlock)
}

func (c *CachedBlockChain) Top() *core.Block {
	l := len(c.cachedBlock)
	if l == 0 {
		return c.BlockChain.Top()
	}
	return c.cachedBlock[l-1]
}
