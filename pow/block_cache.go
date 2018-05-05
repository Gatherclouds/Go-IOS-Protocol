package pow

import (
	"github.com/ethereum/go-ethereum/core"
	"bytes"
)

type CacheStatus int

const (
	Extend     CacheStatus = iota
	Fork
	NotFound
	ErrorBlock
)

type BlockCacheTree struct {
	depth    int
	bc       CachedBlockChain
	children []*BlockCacheTree
	super    *BlockCacheTree
}

func newBct(block *core.Block, tree *BlockCacheTree) *BlockCacheTree {
	bct := BlockCacheTree{
		depth:    0,
		bc:       tree.bc.Copy(),
		children: make([]*BlockCacheTree, 0),
		super:    tree,
	}

	bct.bc.Push(block)
	return &bct
}

func (b *BlockCacheTree) add(block *core.Block, verifier func(blk *core.Block, chain core.BlockChain) bool) CacheStatus {

	var code CacheStatus
	for _, bct := range b.children {
		code = bct.add(block, verifier)
		if code == Extend {
			if bct.depth == b.depth {
				b.depth++
				return Extend
			} else {
				return Fork
			}
		} else if code == Fork {
			return Fork
		} else if code == ErrorBlock {
			return ErrorBlock
		}
	}

	if bytes.Equal(b.bc.Top().Head.Hash(), block.Head.ParentHash) {
		if !verifier(block, &b.bc) {
			return ErrorBlock
		}

		if len(b.children) == 0 {
			b.children = append(b.children, newBct(block, b))
			b.depth++
			return Extend
		} else {
			return Fork
		}
	}

	return NotFound
}

func (b *BlockCacheTree) pop() *BlockCacheTree {

	for _, bct := range b.children {
		if bct.depth == b.depth-1 {
			return bct
		}
	}
	return nil
}

func (b *BlockCacheTree) iterate(fun func(bct *BlockCacheTree) bool) bool {
	if fun(b) {
		return true
	}
	for _, bct := range b.children {
		f := bct.iterate(fun)
		if f {
			return true
		}
	}
	return false
}