package common

import (
	"testing"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ethereum/go-ethereum/core"
)

func TestBlockCache(t *testing.T) {
	b0 := core.Block{
		Head: core.BlockHead{
			ParentHash: []byte("nothing"),
		},
		Content: []byte("b0"),
	}

	b1 := core.Block{
		Head: core.BlockHead{
			ParentHash: b0.HeadHash(),
		},
		Content: []byte("b1"),
	}

	b2 := core.Block{
		Head: core.BlockHead{
			ParentHash: b1.HeadHash(),
		},
		Content: []byte("b2"),
	}

	b2a := core.Block{
		Head: core.BlockHead{
			ParentHash: b1.HeadHash(),
		},
		Content: []byte("fake"),
	}

	b3 := core.Block{
		Head: core.BlockHead{
			ParentHash: b2.HeadHash(),
		},
		Content: []byte("b3"),
	}

	b4 := core.Block{
		Head: core.BlockHead{
			ParentHash: b3.HeadHash(),
		},
	}

	ctl := gomock.NewController(t)

	verifier := func(blk *core.Block, chain core.BlockChain) bool {
		return true
	}

	base := core_mock.NewMockBlockChain(ctl)
	base.EXPECT().Top().AnyTimes().Return(&b0)

	Convey("Test of Block Cache", t, func() {
		Convey("Add:", func() {
			Convey("normal:", func() {
				bc := NewBlockCache(base, 4)
				err := bc.Add(&b1, verifier)
				So(err, ShouldBeNil)
				So(bc.cachedRoot.depth, ShouldEqual, 1)

			})

			Convey("fork and error", func() {
				bc := NewBlockCache(base, 4)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				So(bc.cachedRoot.depth, ShouldEqual, 2)

				verifier = func(blk *core.Block, chain core.BlockChain) bool {
					return false
				}
				err := bc.Add(&b3, verifier)
				So(err, ShouldNotBeNil)
			})

			Convey("auto push", func() {
				var ans string
				base.EXPECT().Push(gomock.Any()).AnyTimes().Do(func(block *core.Block) error {
					ans = string(block.Content)
					return nil
				})
				verifier = func(blk *core.Block, chain core.BlockChain) bool {
					return true
				}
				bc := NewBlockCache(base, 3)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b3, verifier)
				bc.Add(&b4, verifier)
				So(ans, ShouldEqual, "b1")
			})
		})
	})
}

func TestBlockCacheDPoS(t *testing.T) {
	ctl := gomock.NewController(t)
	pool := core_mock.NewMockPool(ctl)

	b0 := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: []byte("nothing"),
			Witness:    "w0",
		},
		Content: []tx.Tx{tx.NewTx(0, nil)},
	}

	b1 := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: b0.HeadHash(),
			Witness:    "w1",
		},
		Content: []tx.Tx{tx.NewTx(1, nil)},
	}

	b2 := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: b1.HeadHash(),
			Witness:    "w2",
		},
		Content: []tx.Tx{tx.NewTx(2, nil)},
	}

	b2a := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: b1.HeadHash(),
			Witness:    "w3",
		},
		Content: []tx.Tx{tx.NewTx(-2, nil)},
	}

	b3 := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: b2.HeadHash(),
			Witness:    "w1",
		},
		Content: []tx.Tx{tx.NewTx(3, nil)},
	}

	b4 := block.Block{
		Head: block.BlockHead{
			Version:    0,
			ParentHash: b2a.HeadHash(),
			Witness:    "w2",
		},
		Content: []tx.Tx{tx.NewTx(4, nil)},
	}

	verifier := func(blk *block.Block, pool state.Pool) (state.Pool, error) {
		return nil, nil
	}

	base := core_mock.NewMockChain(ctl)
	base.EXPECT().Top().AnyTimes().Return(&b0)

	Convey("Test of Block Cache (DPoS)", t, func() {
		Convey("Add:", func() {
			var ans int64
			base.EXPECT().Push(gomock.Any()).Do(func(block *block.Block) error {
				ans = block.Content[0].Nonce
				return nil
			})
			Convey("auto push", func() {
				ans = 0
				bc := NewBlockCache(base, pool, 2)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b3, verifier)
				bc.Add(&b4, verifier)
				So(ans, ShouldEqual, 1)
			})

			Convey("deal with singles", func() {
				ans = 0
				bc := NewBlockCache(base, pool, 2)
				bc.Add(&b2, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b3, verifier)
				bc.Add(&b4, verifier)
				So(len(bc.singleBlockRoot.children), ShouldEqual, 2)
				bc.Add(&b1, verifier)
				So(len(bc.singleBlockRoot.children), ShouldEqual, 0)
				So(ans, ShouldEqual, 1)
			})
		})

		Convey("Longest chain", func() {
			Convey("no forked", func() {
				bc := NewBlockCache(base, pool, 10)
				bc.Add(&b1, verifier)
				bc.Add(&b2, verifier)
				ans := bc.LongestChain().Top().Content[0].Nonce
				So(ans, ShouldEqual, 2)
			})

			Convey("forked", func() {
				var bc BlockCache = NewBlockCache(base, pool, 10)

				bc.Add(&b1, verifier)
				bc.Add(&b2a, verifier)
				bc.Add(&b2, verifier)
				ans := bc.LongestChain().Top().Content[0].Nonce
				So(ans, ShouldEqual, -2)
				bc.Add(&b3, verifier)
				ans = bc.LongestChain().Top().Content[0].Nonce
				So(ans, ShouldEqual, 3)
			})
		})

	})
}