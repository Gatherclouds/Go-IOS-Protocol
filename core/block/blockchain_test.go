package block

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"Go-IOS-Protocol/core/tx"
	"sync"
)

func TestNewBlockChain(t *testing.T) {
	Convey("test TestNewBlockChain", t, func() {
		txDb:=tx.TxDbInstance()
		So(txDb, ShouldNotBeNil)

		bc, err := Instance()
		Convey("New", func() {
			So(err, ShouldBeNil)
			So(bc.Length(), ShouldEqual, bc.Length())
		})
	})
}

