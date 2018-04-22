package core

//go:generate mockgen -destination mocks/mock_statepool.go -package core_mock -source utxo_pool.go -imports .=github.com/iost-official/prototype/core

/*
Current states of system ALERT: 正在施工，请不要调用
*/
type UTXOPool interface {
	Add(state UTXO) error
	Find(stateHash []byte) (UTXO, error)
	Del(StateHash []byte) error
	Transact(block *Block) error
	
}

func BuildStatePoolCore(chain BlockChain) *StatePoolCore {
	var spc StatePoolCore
	spc.cli, _ = redis.Dial(Conn, DBAddr) // TODO : rebuild pool by block chain
	return &spc
}

type StatePoolImpl struct {
	*StatePoolCore
	addList []UTXO
	delList [][]byte
	base    *StatePoolImpl
}