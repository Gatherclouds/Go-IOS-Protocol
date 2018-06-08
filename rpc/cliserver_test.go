package rpc

import (
	"testing"
	"Go-IOS-Protocol/vm/lua"
	"Go-IOS-Protocol/vm"
	"Go-IOS-Protocol/account"

	"github.com/golang/mock/gomock"
	"Go-IOS-Protocol/protocol/mocks"
	"Go-IOS-Protocol/network"
	"context"
	"Go-IOS-Protocol/core/tx"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHttpServer(t *testing.T) {
	Convey("Test of HttpServer", t, func() {
		txDb:=tx.TxDbInstance()
		So(txDb, ShouldNotBeNil)
		main := lua.NewMethod("main", 0, 1)
		code := `function main()
			 		Put("hello", "world")
					return "success"
				end`
		lc := lua.NewContract(vm.ContractInfo{Prefix: "test", GasLimit: 100, Price: 1, Publisher: vm.IOSTAccount("ahaha")}, code, main)

		_tx := tx.NewTx(int64(0), &lc)
		acc, _ := account.NewAccount(nil)
		a1, _ := account.NewAccount(nil)
		sig1, _ := tx.SignContract(_tx, a1)
		_tx, _ = tx.SignTx(_tx, acc, sig1)

		Convey("Test of PublishTx", func() {
			ctl := gomock.NewController(t)
			mockRouter := protocol_mock.NewMockRouter(ctl)
			mockRouter.EXPECT().Broadcast(gomock.Any()).AnyTimes().Return()
			network.Route = mockRouter
			txpb := Transaction{Tx: _tx.Encode()}
			hs := new(HttpServer)
			res, err := hs.PublishTx(context.Background(), &txpb)
			So(err, ShouldBeNil)
			So(res.Code, ShouldEqual, 0)
		})

	})
}
