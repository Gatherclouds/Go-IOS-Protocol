package protocol

import (
	"fmt"
	"github.com/iost-official/Go-IOS-Protocol/iosbase"
)

//go:generate mockgen -destination mocks/mock_database.go -package protocol_mock github.com/iost-official/Go-IOS-Protocol/protocol Database

/*
The runtime database used by protocol
*/
type Database interface {
	NewViewSignal() (chan View, error)

	VerifyTx(tx iosbase.Tx) error
	VerifyTxWithCache(tx iosbase.Tx, cachePool iosbase.TxPool) error
	VerifyBlock(block *iosbase.Block) error
	VerifyBlockWithCache(block *iosbase.Block, cachePool iosbase.TxPool) error

	PushBlock(block *iosbase.Block) error

	GetStatePool() (iosbase.StatePool, error)
	GetBlockChain() (iosbase.BlockChain, error)
	GetCurrentView() (View, error)
}

