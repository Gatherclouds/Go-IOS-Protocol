package core

import "fmt"

type Serializable interface {
	Encode() []byte
	Decode([]byte) error
	Hash() []byte
}

//go:generate mockgen -destination mocks/mock_tx_pool.go -package core_mock github.com/iost-official/prototype/core TxPool
type TxPool interface {
	Add(tx Tx) error
	Del(tx Tx) error
	Find(txHash []byte) (Tx, error)
	GetSlice() ([]Tx, error)
	Has(txHash []byte) (bool, error)
	Size() int
	Serializable

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

func (tp *TxPoolImpl) Find(txHash []byte) (Tx, error) {
	tx, ok := tp.txMap[common.Base58Encode(txHash)]
	if !ok {
		return tx, fmt.Errorf("not found")
	}
	return tx, nil
}

func (tp *TxPoolImpl) Has(txHash []byte) (bool, error) {

	_, ok := tp.txMap[common.Base58Encode(txHash)]
	return ok, nil
}

func (tp *TxPoolImpl) GetSlice() ([]Tx, error) {
	var txs []Tx
	for _, v := range tp.txMap {
		txs = append(txs, v)
	}

	return txs, nil
}

func (tp *TxPoolImpl) Size() int {
	if tp.txMap == nil {
		tp.txMap = make(map[string]Tx)
		return 0
	}
	return len(tp.txMap)
}


