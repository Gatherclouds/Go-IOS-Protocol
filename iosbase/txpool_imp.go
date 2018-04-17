package iosbase

import "fmt"

type TxPoolImpl struct {
	TxPoolRaw
	txMap map[string]Tx
}

