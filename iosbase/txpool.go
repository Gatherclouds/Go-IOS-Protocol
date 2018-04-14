package iosbase

import "fmt"

//go:generate mockgen -destination mocks/mock_txpool.go -package iosbase_mock github.com/iost-official/Go-IOS-Protocol/iosbase TxPool
type TxPool interface {
	Add(tx Tx) error
	Del(tx Tx) error
	Find(txHash []byte) (Tx, error)
	GetSlice() ([]Tx, error)
	Has(txHash []byte) (bool, error)
	Size() int
	Serializable
}

