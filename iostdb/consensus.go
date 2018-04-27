package iostdb

type TxStatus int

type Consensus interface {
	Init(bc BlockChain, sp StatePool, network Net) error
	Run()
	Stop()
	PublishTx(tx Tx) error
	CheckTx(tx Tx) (TxStatus, error)
	GetStatus() (BlockChain, StatePool, error)
	GetCachedStatus() (BlockChain, StatePool, error)
}
