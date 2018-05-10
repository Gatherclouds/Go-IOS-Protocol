package pow

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

}
