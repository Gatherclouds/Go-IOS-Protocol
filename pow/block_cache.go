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

type BlockCache interface {
	Add(block *core.Block, verifier func(blk *core.Block, chain core.BlockChain) bool) error
	FindBlockInCache(hash []byte) (*core.Block, error)
	LongestChain() core.BlockChain
}

type BlockCacheImpl struct {
	bc           core.BlockChain
	cachedRoot   *BlockCacheTree
	singleBlocks []*core.Block
	maxDepth     int
}

func NewBlockCache(chain core.BlockChain, maxDepth int) BlockCacheImpl {
	h := BlockCacheImpl{
		bc: chain,
		cachedRoot: &BlockCacheTree{
			depth:    0,
			bc:       NewCBC(chain),
			children: make([]*BlockCacheTree, 0),
			super:    nil,
		},
		singleBlocks: make([]*core.Block, 0),
		maxDepth:     maxDepth,
	}
	return h
}

func (h *BlockCacheImpl) Add(block *core.Block, verifier func(blk *core.Block, chain core.BlockChain) bool) error {
	code := h.cachedRoot.add(block, verifier)
	switch code {
	case Extend:
		if h.cachedRoot.depth > h.maxDepth {
			h.cachedRoot = h.cachedRoot.pop()
			h.cachedRoot.super = nil
			h.cachedRoot.bc.Flush()
		}
		fallthrough
	case Fork:
		for _, blk := range h.singleBlocks {
			h.Add(blk, verifier)
		}
	case NotFound:
		h.singleBlocks = append(h.singleBlocks, block)
	case ErrorBlock:
		return fmt.Errorf("error found")
	}
	return nil
}