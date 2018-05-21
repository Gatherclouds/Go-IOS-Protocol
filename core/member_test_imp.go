package core

type TxPoolImpl struct {
	TxPoolRaw
	txMap map[string]Tx
}

func NewTxPool() TxPool {
	txp := TxPoolImpl{
		txMap: make(map[string]Tx),
		TxPoolRaw: TxPoolRaw{
			Txs:    make([]Tx, 0),
			TxHash: make([][]byte, 0),
		},
	}
	return &txp
}


