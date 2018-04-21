package core

type Serializable interface {
	Encode() []byte
	Decode([]byte) error
	Hash() []byte
}

//go:generate mockgen -destination mocks/mock_tx_pool.go -package core_mock github.com/iost-official/prototype/core TxPool
type TxPool interface {
	Add(tx Tx) error
	Del(tx Tx) error
	
}

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

func (tp *TxPoolImpl) Add(tx Tx) error {
	tp.txMap[common.Base58Encode(tx.Hash())] = tx
	return nil
}

func (tp *TxPoolImpl) Del(tx Tx) error {
	delete(tp.txMap, common.Base58Encode(tx.Hash()))
	return nil
}

