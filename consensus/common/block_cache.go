package common

import (
	"github.com/ethereum/go-ethereum/core"
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/core/state"
)

//const (
//	MaxCacheDepth = 6
//)

type CacheStatus int

const (
	Extend     CacheStatus = iota
	Fork
	NotFound
	ErrorBlock
)

type BlockCacheTree struct {
	bc       CachedBlockChain
	children []*BlockCacheTree
	super    *BlockCacheTree
	pool     state.Pool
}

func newBct(block *block.Block, tree *BlockCacheTree) *BlockCacheTree {
	bct := BlockCacheTree{
		bc:       tree.bc.Copy(),
		children: make([]*BlockCacheTree, 0),
		super:    tree,
	}

	bct.bc.Push(block)
	return &bct
}

func (b *BlockCacheTree) add(block *block.Block, verifier func(blk *block.Block, pool state.Pool) (state.Pool, error)) (CacheStatus, *BlockCacheTree) {
	for _, bct := range b.children {
		code, newTree := bct.add(block, verifier)
		if code != NotFound {
			return code, newTree
		}
	}

	if bytes.Equal(b.bc.Top().Head.Hash(), block.Head.ParentHash) {
		newPool, err := verifier(block, b.pool)
		if err != nil {
			return ErrorBlock, nil
		}

		bct := newBct(block, b)
		bct.pool = newPool
		b.children = append(b.children, bct)
		if len(b.children) == 1 {
			return Extend, bct
		} else {
			return Fork, bct
		}
	}

	return NotFound, nil
}

func (b *BlockCacheTree) findSingles(block *block.Block) (bool, *BlockCacheTree) {
	for _, bct := range b.children {
		found, ret := bct.findSingles(block)
		if found {
			return found, ret
		}
	}
	if b.bc.block != nil && bytes.Equal(b.bc.block.Head.Hash(), block.Head.ParentHash) {
		return true, b
	}
	return false, nil
}

func (b *BlockCacheTree) addSubTree(root *BlockCacheTree, verifier func(blk *block.Block, pool state.Pool) (state.Pool, error)) {
	block := root.bc.block
	newPool, err := verifier(block, b.pool)
	if err != nil {
		return
	}

	newTree := newBct(block, b)
	newTree.pool = newPool
	b.children = append(b.children, newTree)
	for _, bct := range root.children {
		newTree.addSubTree(bct, verifier)
	}
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

func (h *BlockCacheImpl) FindBlockInCache(hash []byte) (*core.Block, error) {
	var pb *core.Block
	found := h.cachedRoot.iterate(func(bct *BlockCacheTree) bool {
		if bytes.Equal(bct.bc.Top().HeadHash(), hash) {
			pb = bct.bc.Top()
			return true
		} else {
			return false
		}
	})

	if found {
		return pb, nil
	} else {
		return nil, fmt.Errorf("not found")
	}
}