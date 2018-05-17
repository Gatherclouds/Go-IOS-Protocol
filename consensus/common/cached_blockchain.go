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
func (c *CachedBlockChain) Push(block *block.Block) error {
	c.block = block
	c.cachedLength++

	// push的时候更新共识相关变量
	switch block.Head.Version {
	case 0:
		// DPoS
		c.confirmed = 1
		witness := block.Head.Witness
		confirmed := make(map[string]int)
		confirmed[witness] = 1
		cbc := c
		for cbc.parent != nil {
			witness = cbc.parent.Top().Head.Witness
			if _, ok := confirmed[witness]; !ok {
				confirmed[witness] = 0
			}
			confirmed[witness]++
			if len(confirmed) > cbc.parent.confirmed {
				// 如果当前分支有更多的确认，则更新父亲的confirmed
				cbc.parent.confirmed = len(confirmed)
			}
			cbc = cbc.parent
		}
		fallthrough
	case 1:
		// PoW
		c.depth = 0
		cbc := c
		depth := 0
		for cbc.parent != nil {
			depth++
			if depth > cbc.parent.depth {
				cbc.parent.depth = depth
			}
			cbc = cbc.parent
		}
	}

	return nil
}

func (c *CachedBlockChain) Length() uint64 {
	return c.Chain.Length() + uint64(c.cachedLength)
}

func (c *CachedBlockChain) Top() *block.Block {
	if c.cachedLength == 0 {
		return c.Chain.Top()
	}
	return c.block
}

func (c *CachedBlockChain) Copy() CachedBlockChain {
	cbc := CachedBlockChain{
		Chain:        c.Chain,
		parent:       c,
		cachedLength: c.cachedLength,
		confirmed:    0,
	}
	return cbc
}

