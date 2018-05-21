package core

type TxPoolImpl struct {
	TxPoolRaw
	txMap map[string]Tx
}



