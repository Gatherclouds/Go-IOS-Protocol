package core

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTxPoolImpl(t *testing.T) {
	Convey("Test of TxPool", t, func() {
		txp := TxPoolImpl{}
		tx := Tx{
			Version: 1,
			Time:    time.Now().Unix(),
		}
		Convey("Add", func() {
			txp.Add(tx)
			So(len(txp.txMap), ShouldEqual, 1)
		})
		
	})
}
