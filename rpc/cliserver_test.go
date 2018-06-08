package rpc

import (
	"testing"
	"Go-IOS-Protocol/vm/lua"
	"Go-IOS-Protocol/vm"
	"Go-IOS-Protocol/account"

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

	})
}
