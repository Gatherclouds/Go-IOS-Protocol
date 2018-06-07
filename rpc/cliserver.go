package rpc

import (
	"context"
	"fmt"

	"Go-IOS-Protocol/common"
	"Go-IOS-Protocol/core/block"
	"Go-IOS-Protocol/core/tx"
	"Go-IOS-Protocol/network"
	"Go-IOS-Protocol/core/message"
)
//go:generate mockgen -destination mock_rpc/mock_rpc.go -package rpc_mock github.com/iost-official/prototype/rpc CliServer

type BInfo struct {
	Head  block.BlockHead
	txCnt int
}
type HttpServer struct {
}

// newHttpServer 初始Http RPC结构体
func newHttpServer() *HttpServer {
	s := &HttpServer{}
	return s
}

func (s *HttpServer) PublishTx(ctx context.Context, _tx *Transaction) (*Response, error) {
	fmt.Println("PublishTx begin")
	var tx1 tx.Tx
	if _tx == nil {
		return &Response{Code: -1}, fmt.Errorf("argument cannot be nil pointer")
	}
	err := tx1.Decode(_tx.Tx)
	if err != nil {
		return &Response{Code: -1}, err
	}

	err = tx1.VerifySelf() //verify Publisher and Signers
	if err != nil {
		return &Response{Code: -1}, err
	}
	
	router.Broadcast(broadTx)

	go func() {
		Cons := consensus.Cons
		if Cons == nil {
			panic(fmt.Errorf("Consensus is nil"))
		}
		Cons.(*dpos.DPoS).ChTx <- broadTx
		fmt.Println("[rpc.PublishTx]:add tx to TxPool")
	}()
	return &Response{Code: 0}, nil
}

func (s *HttpServer) GetTransaction(ctx context.Context, txkey *TransactionKey) (*Transaction, error) {
	fmt.Println("GetTransaction begin")
	if txkey == nil {
		return nil, fmt.Errorf("argument cannot be nil pointer")
	}
	PubKey := common.Base58Decode(txkey.Publisher)
	//check length of Pubkey here
	if len(PubKey) != 33 {
		return nil, fmt.Errorf("PubKey invalid")
	}
	Nonce := txkey.Nonce
	//check Nonce here

	txDb := tx.TxDb
	if txDb == nil {
		panic(fmt.Errorf("TxDb should be nil"))
	}
	tx, err := txDb.(*tx.TxPoolDb).GetByPN(Nonce, PubKey)
	if err != nil {
		return nil, err
	}

	return &Transaction{Tx: tx.Encode()}, nil
}

